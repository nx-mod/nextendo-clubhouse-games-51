# clubhouse-games-51

Nextendo Network server for **Clubhouse Games: 51 Worldwide Classics** (Nintendo Switch), update 2.0.1.

| | |
|---|---|
| Title ID | `010047700D540000` |
| Game server ID | `2035BB00` (SNI `g2035bb00-lp1.s.n.srv.nintendo.net`) |
| Access key | `8c3661ff` |
| NEX version | unconfirmed (`40605` default) |

The game server ID comes from the hostname the game resolves when entering Online Play. The access
key is not in any public list. It was recovered on 2026-09-15 from one real PRUDP CONNECT
signature captured against a placeholder key: the signature is
`HMAC-MD5(MD5(accessKey), MD5(accessKey)+connectionSig)`, so a search over all 2^32 8-hex-digit
keys found exactly one match, after first confirming the connection signature against the
server's own expected value.

## Servers

- auth `:8460` (TicketGranting, behind the SNI router on `:443`)
- secure `:61001` (SecureConnection, MatchmakeExtension, MatchMaking, MatchMakingExt, NATTraversal, Utility, Ranking)
- dashboard `:8096` (`/api/stats`, `/api/kick`, token-gated)

## Configuration

Copy `example.env` to `.env`. Every unconfirmed wire value is an env var (`CLUBHOUSE_*`), so a wrong
guess costs a restart, not a rebuild. Builds against a sibling `../nextendo-nex` checkout.

## Disclaimer

This server ships no Nintendo code, keys, or copyrighted assets. It is an independent
reimplementation for use with a community-run replacement service, not affiliated with, endorsed
by, or associated with Nintendo.
