package nfsd

import (
	"github.com/mit-pdos/dafny-nfsd/dafny_go/jrnl"
	dirfs "github.com/mit-pdos/dafny-nfsd/dafnygen/DirFs_Compile"
	std "github.com/mit-pdos/dafny-nfsd/dafnygen/Std_Compile"

	"github.com/tchajed/goose/machine/disk"
)

type Nfs struct {
	filesys *dirfs.DirFilesys
}

func zeroDisk(d disk.Disk) {
	zeroblock := make([]byte, 4096)
	sz := d.Size()
	for i := uint64(0); i < sz; i++ {
		d.Write(i, zeroblock)
	}
	d.Barrier()
}

func MakeNfs(d disk.Disk) *Nfs {
	// only runs initialization, recovery isn't set up yet
	zeroDisk(d)
	dfsopt := dirfs.Companion_DirFilesys_.New(&d)
	if dfsopt.Is_None() {
		panic("no dirfs")
	}

	dfs := dfsopt.Get().(std.Option_Some).X.(*dirfs.DirFilesys)

	nfs := &Nfs{
		filesys: dfs,
	}

	return nfs
}

func RecoverNfs(d disk.Disk) *Nfs {
	jrnl := jrnl.NewJrnl(&d)
	dfs := dirfs.New_DirFilesys_()
	dfs.Recover(jrnl)

	nfs := &Nfs{
		filesys: dfs,
	}

	return nfs
}
