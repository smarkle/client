package libkbfs

import (
	"fmt"

	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/protocol/go"
)

type MDOpsConcurTest struct {
	uid   keybase1.UID
	enter chan struct{}
	start chan struct{}
}

func NewMDOpsConcurTest(uid keybase1.UID) *MDOpsConcurTest {
	return &MDOpsConcurTest{
		uid:   uid,
		enter: make(chan struct{}),
		start: make(chan struct{}),
	}
}

func (m *MDOpsConcurTest) GetForHandle(handle *DirHandle) (
	*RootMetadata, error) {
	return nil, fmt.Errorf("Not supported")
}

func (m *MDOpsConcurTest) GetForTLF(id DirID) (*RootMetadata, error) {
	fmt.Printf("HERE in GETFORTLF\n")
	_, ok := <-m.enter
	if !ok {
		// Only one caller should ever get here
		return nil, fmt.Errorf("More than one caller to GetForTLF()!")
	}
	<-m.start
	dh := NewDirHandle()
	dh.Writers = append(dh.Writers, m.uid)
	return NewRootMetadata(dh, id), nil
}

func (m *MDOpsConcurTest) Get(mdID MdID) (*RootMetadata, error) {
	return nil, fmt.Errorf("Not supported")
}

func (m *MDOpsConcurTest) GetSince(id DirID, mdID MdID, max int) (
	[]*RootMetadata, bool, error) {
	return nil, false, nil
}

func (m *MDOpsConcurTest) Put(id DirID, md *RootMetadata, deviceID libkb.KID,
	unmergedBase MdID) error {
	<-m.start
	<-m.enter
	md.SerializedPrivateMetadata = make([]byte, 1, 1)
	return nil
}

func (m *MDOpsConcurTest) PutUnmerged(id DirID, rmd *RootMetadata,
	deviceID libkb.KID) error {
	return nil
}

func (m *MDOpsConcurTest) GetLastCommittedPoint(id DirID, deviceID libkb.KID) (
	bool, MdID, error) {
	return false, MdID{0}, nil
}

func (m *MDOpsConcurTest) GetUnmergedSince(id DirID, deviceID libkb.KID,
	mdID MdID, max int) ([]*RootMetadata, bool, error) {
	return nil, false, nil
}

func (m *MDOpsConcurTest) GetFavorites() ([]DirID, error) {
	return []DirID{}, nil
}
