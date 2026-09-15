package main

// Clubhouse DataStore (0x73) stub. Online account setup calls DataStore before matchmaking; with
// no handler the server returns NotImplemented and the game shows "Communication error."
//
// First-time setup calls GetPersistenceInfos (method 30): ownerId + List<u16> slotIds, response
// List<DataStorePersistenceInfo> + List<Result>. For a new player nothing is saved yet, so we
// answer empty infos and NotFound (0x80690004) per requested slot — the same "nothing saved"
// signal method 8 (GetMeta) uses — which tells the game to create a fresh profile. Every other
// method gets a minimal empty-list response (method 8 also NotFound), matching MHGU/MPS.

import (
	"fmt"

	nex "github.com/NextendoNetwork/nextendo-nex"
)

const (
	protocolDataStore uint16 = 0x73

	methodGetPersistenceInfos uint32 = 30

	dataStoreNotFound uint32 = 0x80690004
)

func dataStoreStubHandler() nex.RMCHandler {
	return func(conn *nex.Connection, req *nex.RMCMessage) *nex.RMCMessage {
		s := conn.Settings

		switch req.Method {
		case methodGetPersistenceInfos:
			in := nex.NewStreamIn(req.Body, s)
			_ = in.U64() // ownerId
			n := in.U32()
			if in.Err() != nil || n > 64 {
				n = 0
			}
			out := nex.NewStreamOut(s)
			out.U32(0) // pPersistenceInfo: empty (no data saved yet)
			out.U32(n) // pResults: one per requested slot
			for i := uint32(0); i < n; i++ {
				out.U32(dataStoreNotFound)
			}
			fmt.Printf("[Clubhouse DataStore] 0x73.30 GetPersistenceInfos slots=%d -> NotFound each (no data yet)\n", n)
			return nex.NewRMCSuccess(s, protocolDataStore, req.Method, req.CallID, out.Bytes())

		case 8: // GetMeta
			fmt.Printf("[Clubhouse DataStore] 0x73.8 -> NotFound 0x80690004\n")
			return nex.NewRMCError(s, protocolDataStore, req.CallID, dataStoreNotFound)

		default:
			out := nex.NewStreamOut(s)
			out.U32(0) // empty list / count = 0
			fmt.Printf("[Clubhouse DataStore] 0x73.%d callID=%d bodyLen=%d -> empty list (stub)\n", req.Method, req.CallID, len(req.Body))
			return nex.NewRMCSuccess(s, protocolDataStore, req.Method, req.CallID, out.Bytes())
		}
	}
}
