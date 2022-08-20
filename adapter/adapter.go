package adapter

type SDCard interface {
	ReadSD() ([]byte, error)
	WriteSD(msg []byte) error
}

type SDCardImpl struct {
}

func (sd *SDCardImpl) ReadSD() ([]byte, error) {
	return []byte{'t', 'e', 's', 't', 's', 'd'}, nil
}

func (sd *SDCardImpl) WriteSD(msg []byte) error {
	return nil
}

type Computer interface {
	ReadSD(card SDCard) ([]byte, error)
}

type ThinkPadComputer struct {
}

func (c *ThinkPadComputer) ReadSD(card SDCard) ([]byte, error) {
	return card.ReadSD()
}

type TFCard interface {
	ReadTF() ([]byte, error)
	WriteTF(msg []byte) error
}

type TFCardImpl struct {
}

func (tf *TFCardImpl) ReadTF() ([]byte, error) {
	return []byte{'t', 'e', 's', 't', 't', 'f'}, nil
}

func (tf *TFCardImpl) WriteTF(msg []byte) error {
	return nil
}

type SDAdapterTF struct {
	tfCard TFCard
}

func (sda *SDAdapterTF) ReadSD() ([]byte, error) {
	return sda.tfCard.ReadTF()
}

func (sda *SDAdapterTF) WriteSD(msg []byte) error {
	return sda.tfCard.WriteTF(msg)
}
