// Copyright 2012-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"net"

	"os"
	"sync/atomic"
	"time"

	"github.com/jacobsa/fuse"
	"github.com/jacobsa/fuse/fuseops"
	"github.com/jacobsa/fuse/fuseutil"
	"github.com/jacobsa/syncutil"
	"harvey-os.org/ninep/protocol"
)

var fid uint32

func newFID() uint32 {
	return sync.AddUint32(&fid, 1)
}

var ino uint64
var newIno() uint64 {
	return sync.AddUint64(&ino)
}

// Create a file system that issues cacheable responses according to the
// following rules:
//
//   - LookUpInodeResponse.Entry.EntryExpiration is set according to
//     lookupEntryTimeout.
//
//   - GetInodeAttributesResponse.AttributesExpiration is set according to
//     getattrTimeout.
//
//   - Nothing else is marked cacheable. (In particular, the attributes
//     returned by LookUpInode are not cacheable.)
func NewP9FS(conn net.Conn, root string, lookupEntryTimeout time.Duration, getattrTimeout time.Duration) (fuse.Server, *P9FS, error) {
	v("attach %v", conn)
	c, err := protocol.NewClient(func(c *protocol.Client) error {
		c.FromNet, c.ToNet = conn, conn
		return nil
	},
		func(c *protocol.Client) error {
			c.Msize = 8192
			c.Trace = v
			return nil
		})
	if err != nil {
		return nil, nil, fmt.Errorf("%v", err)
	}
	msize, vers, err := c.CallTversion(8000, "9P2000")
	if err != nil {
		return nil, nil, fmt.Errorf("CallTversion: want nil, got %v", err)
	}
	v("CallTversion: msize %v version %v", msize, vers)
	if _, err := c.CallTattach(0, protocol.NOFID, "", root); err != nil {
		return nil, nil, err
	}

	cfs := &P9FS{
		cl:                 c,
		lookupEntryTimeout: lookupEntryTimeout,
		getattrTimeout:     getattrTimeout,
		mtime:              time.Now(),
		inMap:              make(map[fuseops.InodeID]entry),
		openfile:           make(map[fuseops.HandleID]openfile),
		ino:                1,
		keepPageCache:      true,
	}

	cfs.inMap[1] = entry{
		fid:      root,
		QID:      p9.QID{Path: 1},
		root:     true,
		fullPath: "/",
		refcount: 1,
	}
	return fuseutil.NewFileSystemServer(cfs), cfs, nil
}

type entry struct {
	file      protocol.File
	fid protocol.FID
	root     bool
	QID      protocol.QID
	fullPath string
	ino      uint64
	refcount uint64
}

type openfile struct {
	fid  protocol.File
	unit int
}

type P9FS struct {
	/////////////////////////
	// Constant data
	/////////////////////////

	lookupEntryTimeout time.Duration
	getattrTimeout     time.Duration
	cl                 *protocol.Client

	/////////////////////////
	// Mutable state
	/////////////////////////
	// unique inumber for this mount's lifetime.
	ino uint64

	mu syncutil.InvariantMutex

	// GUARDED_BY(mu)
	keepPageCache bool
	mtime         time.Time
	inMap         map[fuseops.InodeID]entry
	openfile      map[fuseops.HandleID]openfile
}

var _ fuseutil.FileSystem = &P9FS{}

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

// LOCKS_REQUIRED(fs.mu)
func (fs *P9FS) rootAttrs() fuseops.InodeAttributes {
	return fuseops.InodeAttributes{
		Mode:  os.ModeDir | 0777,
		Mtime: fs.mtime,
	}
}

////////////////////////////////////////////////////////////////////////
// Public interface
////////////////////////////////////////////////////////////////////////

// LOCKS_EXCLUDED(fs.mu)
func (fs *P9FS) SetMtime(mtime time.Time) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	fs.mtime = mtime
}

// LOCKS_EXCLUDED(fs.mu)
func (fs *P9FS) SetKeepCache(keep bool) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	fs.keepPageCache = keep
}

////////////////////////////////////////////////////////////////////////
// FileSystem methods
////////////////////////////////////////////////////////////////////////

func (fs *P9FS) StatFS(ctx context.Context, op *fuseops.StatFSOp) error {
	return nil
}

// LookupInode looks up an inode from an inode FUSE already looked up.
// There's a subtle point here: will FUSE ask us to look up an inode it knows about.
// The answer seems to be no. Therefore, we can dynamically generate an inumber
// if we get here and not worry that there's an "obsolete" inumber for a previous
// lookup. E.g., if fuse asks us to lookup /dev, and some time later it asks us
// again, it's because it told us to forget about it and there's no in-kernel
// record of it. IOW, we take it as given that the kernel gets it right.
// This is kind of ok, since the kernel has a much longer lifetime than an
// instance of this client, which after all is just one cpu session.
// Hence, we don't use the QID.Path, as it is not unique anyway (unlike Plan 9)
// but, rather, we dynamically generate a new inumber for each lookup that
// succeeds.
// LOCKS_EXCLUDED(fs.mu)
func (p9fs *P9FS) LookUpInode(ctx context.Context, op *fuseops.LookUpInodeOp) error {
	p9fs.mu.Lock()
	defer p9fs.mu.Unlock()

	// Find the ID and attributes.
	p := op.Parent
	cl, ok := p9fs.inMap[p]
	if !ok {
		panic("NO parent")
		return os.ErrNotExist
	}

	w, err := cl.CallTwalk([]string{op.Name})

	if err != nil {
		//log.Panicf("walkgetattr: %v walking from %v in %d", err, cl, p)
		return err
	}

	q := qids[0]
	ino := atomic.AddUint64(&p9fs.ino, 1)
	i, ok := p9fs.inMap[fuseops.InodeID(ino)]
	if ok {
		log.Panicf("WTF? lookup %v and inumber %d was taken?", f, i)
	}
	log.Printf("CPUD: at inmap: %v", i)

	p9fs.inMap[fuseops.InodeID(ino)] = entry{
		fid:  f,
		root: false,

		QID:  q,

		ino:  ino,
	}
	/*
		Mode             FileMode
		UID              UID
		GID              GID
		NLink            NLink
		RDev             Dev
		Size             uint64
		BlockSize        uint64
		Blocks           uint64
		ATimeSeconds     uint64
		ATimeNanoSeconds uint64
		MTimeSeconds     uint64
		MTimeNanoSeconds uint64
		CTimeSeconds     uint64
		CTimeNanoSeconds uint64
		BTimeSeconds     uint64
		BTimeNanoSeconds uint64
		Gen              uint64
		DataVersion      uint64
	*/
	var dir fs.FileMode
	if q.Type&protocol.QTDIR == protocol.QTDIR {
		dir = os.ModeDir
	}
	//	var dt = ptype(q)
	attrs := fuseops.InodeAttributes{
		Size:  a.Size,
		Nlink: uint32(a.NLink),
		Mode:  dir | fs.FileMode(a.Mode),
		Atime: time.Unix(int64(a.ATimeSeconds), int64(a.ATimeNanoSeconds)),
		Mtime: time.Unix(int64(a.MTimeSeconds), int64(a.MTimeNanoSeconds)),
		Ctime: time.Unix(int64(a.CTimeSeconds), int64(a.CTimeNanoSeconds)),
		Uid:   uint32(a.UID),
		Gid:   uint32(a.GID),
	}
	v("attrs %#x", attrs)
	// Fill in the response.
	op.Entry.Child = fuseops.InodeID(ino)
	op.Entry.Attributes = attrs
	op.Entry.EntryExpiration = time.Now().Add(p9fs.lookupEntryTimeout)

	return nil
}

func ptype(q protocol.QID) fuseutil.DirentType {
	/*	DT_Unknown   DirentType = 0
		DT_Socket    DirentType = syscall.DT_SOCK
		DT_Link      DirentType = syscall.DT_LNK
		DT_File      DirentType = syscall.DT_REG
		DT_Block     DirentType = syscall.DT_BLK
		DT_Directory DirentType = syscall.DT_DIR
		DT_Char      DirentType = syscall.DT_CHR
		DT_FIFO      DirentType = syscall.DT_FIFO
	*/
	switch {
	case q.Type&protocolQTDIR == protocol.QTDIR:
		return fuseutil.DT_Directory
		//	case q.Type.IsSocket(), q.Type.IsNamedPipe(), q.Type.IsCharacterDevice():
		// Best approximation.
		//		return fuseutil.DT_Socket
		//	case q.Type.IsSymlink():
		//		return fuseutil.DT_Link
	default:
		return fuseutil.DT_File
	}
}

// LOCKS_EXCLUDED(fs.mu)
func (p9fs *P9FS) GetInodeAttributes(ctx context.Context, op *fuseops.GetInodeAttributesOp) error {
	p9fs.mu.Lock()
	defer p9fs.mu.Unlock()

	// Figure out which inode the request is for.
	in := op.Inode
	cl, ok := p9fs.inMap[in]
	if !ok {
		panic("NO file")
		return os.ErrNotExist
	}

	return errors.New("not yet")
}

// OpenDir implements OpenDir. N.B.: need to do a walk and open,
// else walks from the directory are impossible!
func (p9fs *P9FS) OpenDir(ctx context.Context, op *fuseops.OpenDirOp) error {
	in := op.Inode
	cl, ok := p9fs.inMap[in]
	if !ok {
		panic("NO file")
		return os.ErrNotExist
	}
	
	q, iounit, err := root.CallTopen(fid, 0)
	if err != nil {
		return nil, err
	}
	if err != nil {
		panic("opendir open")
		return err
	}
	h := newIno()
	op.Handle = fuseops.HandleID(h)

	p9fs.openfile[op.Handle] = openfile{
		file:  f,
		FID: fid,
		unit: int(unit),
	}

	return nil
}

func (fs *P9FS) ReadDir(ctx context.Context, op *fuseops.ReadDirOp) error {
	ha := op.Handle
	cl, ok := fs.openfile[ha]
	if !ok {
		panic("NO open file")
		return os.ErrNotExist
	}

	// The offset is determined by the rather arbitrary value from 9p.
	off := op.Offset

	d, err := cl.fid.Readdir(uint64(off), uint32(cl.unit))
	if err != nil {
		panic("NO readdir")
		return err
	}

	var tot int
	for _, ent := range d {
		// you get QID, Offset, Type, and Name.
		/*	DT_Unknown   DirentType = 0
			DT_Socket    DirentType = syscall.DT_SOCK
			DT_Link      DirentType = syscall.DT_LNK
			DT_File      DirentType = syscall.DT_REG
			DT_Block     DirentType = syscall.DT_BLK
			DT_Directory DirentType = syscall.DT_DIR
			DT_Char      DirentType = syscall.DT_CHR
			DT_FIFO      DirentType = syscall.DT_FIFO
		*/
		var dt = ptype(ent.QID)

		fe := fuseutil.Dirent{
			Offset: fuseops.DirOffset(ent.Offset),
			Inode:  fuseops.InodeID(ent.QID.Path),
			Name:   ent.Name,
			Type:   dt,
			Inode:  fuseops.InodeID(ent.QID.Path),
		}
		n := fuseutil.WriteDirent(op.Dst[tot:], fe)
		tot += n
	}
	op.BytesRead = tot

	return nil
}

// OpenFile implements OpenFile.
// Again, we take FUSE as authoritative: if it is asking us to open a file, it is because
// it needs it opened.
func (p9fs *P9FS) OpenFile(ctx context.Context, op *fuseops.OpenFileOp) error {
	p9fs.mu.Lock()
	defer p9fs.mu.Unlock()

	in := op.Inode
	cl, ok := p9fs.inMap[in]
	if !ok {
		panic("NO file")
		return os.ErrNotExist
	}

	fid := newFID()
	// We walk because it is allowed to walk a file fid to another fid.
	// Were we to open this fid, it would be breaking the rules.
	_, err := cl.CallTWalk(c.FID, fid, []string{})
	if err != nil {
		panic("openfile walk")
		return err
	}
	q, iouint, err := cl.CallTOpen(fid, 0)
	if err != nil {
		panic("openfile open")
		return err
	}

	h := atomic.AddUint64(&p9fs.ino, 1)
	op.Handle = fuseops.HandleID(h)

	p9fs.openfile[op.Handle] = openfile{
		fid:  f,
		unit: int(unit),
	}

=======
	/*
		// We walk because it is allowed to walk a file fid to another fid.
		// Were we to open this fid, it would be breaking the rules.
		_, f, err := cl.fid.Walk([]string{})
		if err != nil {
			panic("openfile walk")
			return err
		}
		_, unit, err := f.Open(ninep.ReadOnly)
		if err != nil {
			panic("openfile open")
			return err
		}

		h := atomic.AddUint64(&p9fs.ino, 1)
		op.Handle = fuseops.HandleID(h)

		p9fs.openfile[op.Handle] = openfile{
			fid:  f,
			unit: int(unit),
		}

	*/
>>>>>>> Stashed changes
	op.KeepPageCache = p9fs.keepPageCache
	return nil
}

// ReadFile implements ReadFile
func (fs *P9FS) ReadFile(ctx context.Context, op *fuseops.ReadFileOp) error {
	ha := op.Handle
	cl, ok := fs.openfile[ha]
	if !ok {
		panic("NO open file")
		return os.ErrNotExist
	}

	off := op.Offset

	dst := op.Dst
	if dst == nil {
		dst = make([]byte, op.Size)
		op.Data = [][]byte{dst}
	}
	amt, err := cl.fid.ReadAt(dst, off)
	op.BytesRead = amt

	return err
}

// The fuse package says to embed a fuseutil.NotImplementedFileSystem in your struct
// to catch all the stuff you don't implement. That way lies madness, we've tried
// it, it's basically undebuggable. So we put all these not implemented bits here.
// A FileSystem that responds to all ops with fuse.ENOSYS. Embed this in your
// struct to inherit default implementations for the methods you don't care
// about, ensuring your struct will continue to implement FileSystem even as
// new methods are added.
func (fs *P9FS) SetInodeAttributes(ctx context.Context, op *fuseops.SetInodeAttributesOp) error {
	panic("func (fs *P9FS) SetInodeAttributes(ctx context.Context, op *fuseops.SetInodeAttributesOp) error {")
	return fuse.ENOSYS
}

// Forget: do we close it too?
func (fs *P9FS) ForgetInode(ctx context.Context, op *fuseops.ForgetInodeOp) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	in := op.Inode
	f, ok := fs.inMap[in]
	if !ok {
		return os.ErrNotExist
	}
	f.refcount -= op.N
	if f.refcount < 1 {
		delete(fs.inMap, in)
	}
	return nil
}

// BatchForget implements BatchForget. Error trees/chains will wait for Go 1.20
func (fs *P9FS) BatchForget(ctx context.Context, op *fuseops.BatchForgetOp) error {
	for i, e := range op.Entries {
		fe := &fuseops.ForgetInodeOp{
			Inode:     e.Inode,
			N:         e.N,
			OpContext: op.OpContext,
		}
		if err := fs.ForgetInode(ctx, fe); err != nil {
			log.Printf("batchforget, entry %d (%v): %v", i, e, err)
		}
	}
	return nil
}

func (fs *P9FS) MkDir(ctx context.Context, op *fuseops.MkDirOp) error {
	panic("func (fs *P9FS) MkDir(ctx context.Context, op *fuseops.MkDirOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) MkNode(ctx context.Context, op *fuseops.MkNodeOp) error {
	panic("func (fs *P9FS) MkNode(ctx context.Context, op *fuseops.MkNodeOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) CreateFile(ctx context.Context, op *fuseops.CreateFileOp) error {
	panic("func (fs *P9FS) CreateFile(ctx context.Context, op *fuseops.CreateFileOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) CreateSymlink(ctx context.Context, op *fuseops.CreateSymlinkOp) error {
	panic("func (fs *P9FS) CreateSymlink(ctx context.Context, op *fuseops.CreateSymlinkOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) CreateLink(ctx context.Context, op *fuseops.CreateLinkOp) error {
	panic("func (fs *P9FS) CreateLink(ctx context.Context, op *fuseops.CreateLinkOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) Rename(ctx context.Context, op *fuseops.RenameOp) error {
	panic("func (fs *P9FS) Rename(ctx context.Context, op *fuseops.RenameOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) RmDir(ctx context.Context, op *fuseops.RmDirOp) error {
	panic("func (fs *P9FS) RmDir(ctx context.Context, op *fuseops.RmDirOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) Unlink(ctx context.Context, op *fuseops.UnlinkOp) error {
	panic("func (fs *P9FS) Unlink(ctx context.Context, op *fuseops.UnlinkOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) ReleaseDirHandle(ctx context.Context, op *fuseops.ReleaseDirHandleOp) error {
	ha := op.Handle
	cl, ok := fs.openfile[ha]
	if !ok {
		return nil
	}
	delete(fs.openfile, ha)
	return cl.fid.Close()
}

func (fs *P9FS) WriteFile(ctx context.Context, op *fuseops.WriteFileOp) error {
	panic("func (fs *P9FS) WriteFile(ctx context.Context, op *fuseops.WriteFileOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) SyncFile(ctx context.Context, op *fuseops.SyncFileOp) error {
	panic("func (fs *P9FS) SyncFile(ctx context.Context, op *fuseops.SyncFileOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) FlushFile(ctx context.Context, op *fuseops.FlushFileOp) error {
	log.Printf("TODO:func (fs *P9FS) FlushFile(ctx context.Context, op *fuseops.FlushFileOp) error {")
	return nil
}

func (fs *P9FS) ReleaseFileHandle(ctx context.Context, op *fuseops.ReleaseFileHandleOp) error {
	ha := op.Handle
	cl, ok := fs.openfile[ha]
	if !ok {
		return nil
	}
	delete(fs.openfile, ha)
	return cl.fid.Close()
}

func (fs *P9FS) ReadSymlink(ctx context.Context, op *fuseops.ReadSymlinkOp) error {
	panic("func (fs *P9FS) ReadSymlink(ctx context.Context, op *fuseops.ReadSymlinkOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) RemoveXattr(ctx context.Context, op *fuseops.RemoveXattrOp) error {
	panic("func (fs *P9FS) RemoveXattr(ctx context.Context, op *fuseops.RemoveXattrOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) GetXattr(ctx context.Context, op *fuseops.GetXattrOp) error {
	v("FIX func (fs *P9FS) GetXattr(ctx context.Context, op *fuseops.GetXattrOp) error {")
	return fuse.ENOATTR
}

func (fs *P9FS) ListXattr(ctx context.Context, op *fuseops.ListXattrOp) error {
	panic("func (fs *P9FS) ListXattr(ctx context.Context, op *fuseops.ListXattrOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) SetXattr(ctx context.Context, op *fuseops.SetXattrOp) error {
	panic("func (fs *P9FS) SetXattr(ctx context.Context, op *fuseops.SetXattrOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) Fallocate(ctx context.Context, op *fuseops.FallocateOp) error {
	panic("func (fs *P9FS) Fallocate(ctx context.Context, op *fuseops.FallocateOp) error {")
	return fuse.ENOSYS
}

func (fs *P9FS) Destroy() {
	panic("func (fs *P9FS) Destroy() {")
}
