package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/injoyai/base/chans"
	"github.com/injoyai/conv"
	"github.com/injoyai/goutil/oss"
	"github.com/injoyai/goutil/oss/tray"
	"github.com/injoyai/proxy/core"
	"github.com/injoyai/proxy/forward"
	"github.com/injoyai/tool/config"
)

func Run(f *forward.Forward) {

	tray.Run(
		func(s *tray.Tray) {

			r := chans.NewRerun(func(ctx context.Context) {
				s.SetHint(fmt.Sprintf("状态: 运行中\n端口: %s\n地址: %s", f.Listen.Port, f.Forward.Address))
				err := f.Run(ctx)
				s.SetHint(fmt.Sprintf("状态: %v\n端口: %s\n地址: %s", err, f.Listen.Port, f.Forward.Address))
			})

			r.Rerun()

			s.AddMenu().SetName("配置").OnClick(func(m *tray.Menu) {
				config.GUI(config.New(oss.UserInjoyDir("/forward/config/config.yaml"), config.Natures{
					{Name: "监听端口", Key: "port", Type: "string"},
					{Name: "转发地址", Key: "address", Type: "string"},
				}).SetWidthHeight(720, 345).OnSaved(func(m *conv.Map) {
					f.Listen = core.NewListenTCP(m.GetInt("port"))
					f.Forward = core.NewDialTCP(m.GetString("address"))
					r.Rerun()
				}))
			})

		},
		tray.WithIco(ico),
		tray.WithLabel("v1.0.1"),
		tray.WithStartup(),
		tray.WithExit(),
	)
}

var ico = []byte{
	0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x20, 0x20, 0x00, 0x00, 0x01, 0x00,
	0x20, 0x00, 0xA8, 0x10, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x28, 0x00,
	0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x01, 0x00,
	0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x76, 0xA7, 0x14, 0x1A, 0x7B, 0xAD,
	0x19, 0x66, 0x79, 0xAB, 0x17, 0xA4, 0x7A, 0xAC, 0x18, 0xD1, 0x7A, 0xAC,
	0x17, 0xEF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x17, 0xEF, 0x7A, 0xAC, 0x18, 0xD1, 0x79, 0xAB, 0x17, 0xA4, 0x7B, 0xAD,
	0x19, 0x66, 0x76, 0xA7, 0x14, 0x1A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7B, 0xAD, 0x17, 0x38, 0x79, 0xAC,
	0x18, 0xA8, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x79, 0xAC, 0x18, 0xA8, 0x7B, 0xAD, 0x17, 0x38, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x80, 0x00, 0x04, 0x7B, 0xAC,
	0x18, 0x93, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7B, 0xAC, 0x18, 0x93, 0x80, 0x80,
	0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7A, 0xAC,
	0x16, 0x2E, 0x79, 0xAC, 0x19, 0xD0, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x79, 0xAC, 0x19, 0xD0, 0x7A, 0xAC,
	0x16, 0x2E, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x7B, 0xAA, 0x1A, 0x3C, 0x7A, 0xAC, 0x18, 0xF0, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xF0, 0x7B, 0xAA,
	0x1A, 0x3C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7A, 0xAC, 0x16, 0x2E, 0x7A, 0xAC,
	0x18, 0xF0, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xF0, 0x7A, 0xAC,
	0x16, 0x2E, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x80,
	0x00, 0x04, 0x79, 0xAC, 0x19, 0xD0, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x79, 0xAC, 0x19, 0xD0, 0x80, 0x80,
	0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7B, 0xAC, 0x18, 0x93, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7B, 0xAC, 0x18, 0x93, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7B, 0xAD,
	0x17, 0x38, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7B, 0xAD, 0x17, 0x38, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x79, 0xAC, 0x18, 0xA8, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x79, 0xAC, 0x18, 0xA8, 0x00, 0x00, 0x00, 0x00, 0x76, 0xA7,
	0x14, 0x1A, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x9B, 0xC1, 0x52, 0xFF, 0xBD, 0xD6, 0x8C, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0xBD, 0xD6, 0x8C, 0xFF, 0xCD, 0xE0, 0xA8, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x76, 0xA7, 0x14, 0x1A, 0x7B, 0xAD, 0x19, 0x66, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0xAC, 0xCB,
	0x6F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xCD, 0xE0, 0xA8, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xEE, 0xF5, 0xE2, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7B, 0xAD,
	0x19, 0x66, 0x79, 0xAB, 0x17, 0xA4, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0xAC, 0xCB, 0x6F, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xDE, 0xEA, 0xC5, 0xFF, 0x8B, 0xB6,
	0x35, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xEE, 0xF5, 0xE2, 0xFF, 0x9B, 0xC1,
	0x52, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x79, 0xAB, 0x17, 0xA4, 0x7A, 0xAC,
	0x18, 0xD1, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xEE, 0xF5,
	0xE2, 0xFF, 0xBD, 0xD6, 0x8C, 0xFF, 0xAC, 0xCB, 0x6F, 0xFF, 0x9B, 0xC1,
	0x52, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xEE, 0xF5, 0xE2, 0xFF, 0x9B, 0xC1,
	0x52, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xD1, 0x7A, 0xAC, 0x17, 0xEF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0xDE, 0xEA, 0xC5, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x9B, 0xC1,
	0x52, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x17, 0xEF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x8B, 0xB6,
	0x35, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xAC, 0xCB,
	0x6F, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x9B, 0xC1,
	0x52, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x8B, 0xB6,
	0x35, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x17, 0xEF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x9B, 0xC1,
	0x52, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xBD, 0xD6, 0x8C, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x17, 0xEF, 0x7A, 0xAC, 0x18, 0xD1, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x8B, 0xB6,
	0x35, 0xFF, 0xCD, 0xE0, 0xA8, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xBD, 0xD6, 0x8C, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xD1, 0x79, 0xAB,
	0x17, 0xA4, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x8B, 0xB6, 0x35, 0xFF, 0xAC, 0xCB, 0x6F, 0xFF, 0xBD, 0xD6,
	0x8C, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xBD, 0xD6,
	0x8C, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x79, 0xAB, 0x17, 0xA4, 0x7B, 0xAD, 0x19, 0x66, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xEE, 0xF5,
	0xE2, 0xFF, 0xBD, 0xD6, 0x8C, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7B, 0xAD,
	0x19, 0x66, 0x76, 0xA7, 0x14, 0x1A, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xEE, 0xF5, 0xE2, 0xFF, 0xBD, 0xD6, 0x8C, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x76, 0xA7, 0x14, 0x1A, 0x00, 0x00,
	0x00, 0x00, 0x79, 0xAC, 0x18, 0xA8, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xEE, 0xF5, 0xE2, 0xFF, 0xBD, 0xD6,
	0x8C, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x79, 0xAC,
	0x18, 0xA8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7B, 0xAD,
	0x17, 0x38, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x8B, 0xB6,
	0x35, 0xFF, 0x9B, 0xC1, 0x52, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7B, 0xAD, 0x17, 0x38, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7B, 0xAC,
	0x18, 0x93, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7B, 0xAC,
	0x18, 0x93, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x80, 0x00, 0x04, 0x79, 0xAC,
	0x19, 0xD0, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x79, 0xAC, 0x19, 0xD0, 0x80, 0x80, 0x00, 0x04, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7A, 0xAC, 0x16, 0x2E, 0x7A, 0xAC,
	0x18, 0xF0, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xF0, 0x7A, 0xAC,
	0x16, 0x2E, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7B, 0xAA, 0x1A, 0x3C, 0x7A, 0xAC,
	0x18, 0xF0, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xF0, 0x7B, 0xAA, 0x1A, 0x3C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7A, 0xAC, 0x16, 0x2E, 0x79, 0xAC,
	0x19, 0xD0, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x79, 0xAC, 0x19, 0xD0, 0x7A, 0xAC, 0x16, 0x2E, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x80, 0x00, 0x04, 0x7B, 0xAC,
	0x18, 0x93, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7B, 0xAC, 0x18, 0x93, 0x80, 0x80,
	0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7B, 0xAD,
	0x17, 0x38, 0x79, 0xAC, 0x18, 0xA8, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x79, 0xAC, 0x18, 0xA8, 0x7B, 0xAD,
	0x17, 0x38, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x76, 0xA7, 0x14, 0x1A, 0x7B, 0xAD, 0x19, 0x66, 0x79, 0xAB,
	0x17, 0xA4, 0x7A, 0xAC, 0x18, 0xD1, 0x7A, 0xAC, 0x17, 0xEF, 0x7A, 0xAC,
	0x18, 0xFF, 0x7A, 0xAC, 0x18, 0xFF, 0x7A, 0xAC, 0x17, 0xEF, 0x7A, 0xAC,
	0x18, 0xD1, 0x79, 0xAB, 0x17, 0xA4, 0x7B, 0xAD, 0x19, 0x66, 0x76, 0xA7,
	0x14, 0x1A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xC0, 0x03, 0xFF, 0xFF, 0x00,
	0x00, 0xFF, 0xFC, 0x00, 0x00, 0x3F, 0xF8, 0x00, 0x00, 0x1F, 0xF0, 0x00,
	0x00, 0x0F, 0xE0, 0x00, 0x00, 0x07, 0xC0, 0x00, 0x00, 0x03, 0xC0, 0x00,
	0x00, 0x03, 0x80, 0x00, 0x00, 0x01, 0x80, 0x00, 0x00, 0x01, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00,
	0x00, 0x01, 0x80, 0x00, 0x00, 0x01, 0xC0, 0x00, 0x00, 0x03, 0xC0, 0x00,
	0x00, 0x03, 0xE0, 0x00, 0x00, 0x07, 0xF0, 0x00, 0x00, 0x0F, 0xF8, 0x00,
	0x00, 0x1F, 0xFC, 0x00, 0x00, 0x3F, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0xC0,
	0x03, 0xFF,
}
