
# Configuration

* server (string) - Hostname/port to connect to. If set, bdn9 will act like a client and send requests to this server.
* bind (string) - IP address/port to bind to and listen for incoming requests. Defaults to "localhost:17432".
* slackWindowName (string) - Name of the slack call window process. Used for mute toggle.
* slackMuteKeys (string) - List of keys to send to slack on toggle mute events. The first entry must be a key and the remaining can be modifier keys.
* teamsWindowName (string) - Name of the teams call window process. Used for mute toggle.
* teamsMuteKeys (string) - List of keys to send to teams on toggle mute events. The first entry must be a key and the remaining can be modifier keys.
