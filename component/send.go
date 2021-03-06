package component

import (
	"github.com/bdlm/log"
	"gosrc.io/xmpp/stanza"
)

func (c *Config) sender(packets chan stanza.Packet) {
	for packet := range packets {
		if p := c.sending(packet); p != nil {
			c.xmpp.Send(p)
		}
	}
}

func (c *Config) sending(packet stanza.Packet) stanza.Packet {
	logger := log.WithField("type", c.Type)
	switch p := packet.(type) {
	case stanza.Message:
		if p.From == "" {
			p.From = c.Host
		} else {
			p.From += "@" + c.Host
		}
		if c.XMPPDebug {
			logger.WithFields(map[string]interface{}{
				"from": p.From,
				"to":   p.To,
				"id":   p.Id,
			}).Debug(p.XMPPFormat())
		}
		return p
	default:
		log.Warn("ignoring packet:", packet)
		return nil
	}
}
