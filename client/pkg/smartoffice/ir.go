package smartoffice

type IRCONTROL_TYPE uint16

const (
	IRCONTROL_TYPE_BangOlufsen            = IRCONTROL_TYPE(1)
	IRCONTROL_TYPE_BangOlufsenDataLink    = IRCONTROL_TYPE(2)
	IRCONTROL_TYPE_BangOlufsenRaw         = IRCONTROL_TYPE(3)
	IRCONTROL_TYPE_BangOlufsenRawDataLink = IRCONTROL_TYPE(4)
	IRCONTROL_TYPE_BoseWave               = IRCONTROL_TYPE(5)
	IRCONTROL_TYPE_Denon                  = IRCONTROL_TYPE(6)
	IRCONTROL_TYPE_DenonRaw               = IRCONTROL_TYPE(7)
	IRCONTROL_TYPE_FAST                   = IRCONTROL_TYPE(8)
	IRCONTROL_TYPE_JVC                    = IRCONTROL_TYPE(9)
	IRCONTROL_TYPE_LG2Repeat              = IRCONTROL_TYPE(10)
	IRCONTROL_TYPE_LG                     = IRCONTROL_TYPE(11)
	IRCONTROL_TYPE_LG2                    = IRCONTROL_TYPE(12)
	IRCONTROL_TYPE_LGRaw                  = IRCONTROL_TYPE(13)
	IRCONTROL_TYPE_NECRepeat              = IRCONTROL_TYPE(14)
	IRCONTROL_TYPE_NEC                    = IRCONTROL_TYPE(15)
	IRCONTROL_TYPE_NEC2                   = IRCONTROL_TYPE(16)
	IRCONTROL_TYPE_NECRaw                 = IRCONTROL_TYPE(17)
	IRCONTROL_TYPE_Onkyo                  = IRCONTROL_TYPE(18)
	IRCONTROL_TYPE_Apple                  = IRCONTROL_TYPE(19)
	IRCONTROL_TYPE_Kaseikyo               = IRCONTROL_TYPE(20)
	IRCONTROL_TYPE_Panasonic              = IRCONTROL_TYPE(21)
	IRCONTROL_TYPE_Kaseikyo_Denon         = IRCONTROL_TYPE(22)
	IRCONTROL_TYPE_Kaseikyo_Mitsubishi    = IRCONTROL_TYPE(23)
	IRCONTROL_TYPE_Kaseikyo_Sharp         = IRCONTROL_TYPE(24)
	IRCONTROL_TYPE_Kaseikyo_JVC           = IRCONTROL_TYPE(25)
	IRCONTROL_TYPE_RC5                    = IRCONTROL_TYPE(26)
	IRCONTROL_TYPE_RC6                    = IRCONTROL_TYPE(27)
	IRCONTROL_TYPE_SamsungLGRepeat        = IRCONTROL_TYPE(28)
	IRCONTROL_TYPE_Samsung                = IRCONTROL_TYPE(29)
	IRCONTROL_TYPE_Samsung48              = IRCONTROL_TYPE(30)
	IRCONTROL_TYPE_SamsungLG              = IRCONTROL_TYPE(31)
	IRCONTROL_TYPE_Sharp                  = IRCONTROL_TYPE(32)
	IRCONTROL_TYPE_Sony                   = IRCONTROL_TYPE(33)
	IRCONTROL_TYPE_LegoPowerFunctions     = IRCONTROL_TYPE(34)
	IRCONTROL_TYPE_MagiQuest              = IRCONTROL_TYPE(35)
)

func newIRSend(deviceType IRCONTROL_TYPE, addressOrHeader, data uint64, nBits uint8, repeats uint8, extra_arg uint64) []byte {
	var cmds Commands
	cmds.SendIR(deviceType, addressOrHeader, data, nBits, repeats, extra_arg)
	return cmds.Bytes()
}
