# Trofiebot/Emotemon

IRC chat bot for Twitch.tv written in Go. Make a Pokemon Go'ish game out of the Twitch emotes.

Features:
- [x] Spam to capture
- [ ] Battle Support
- [ ] User's Emotemon

# Setup

1. Clone the repo
2. Configure the bot:
  1. Edit `bot/bot.go`:
    - Add your bots name (*Twitch Name*)
    - Add the desired Twitch chat channel
  2. Make a `twitch_pass.txt` file in the **$GOPATH/bin** folder
  3. Place your Twitch OAUTH token (corresponding to the bot's Twitch name) within `twitch_pass.txt`
3. Run `go install` trofiebot
4. Log on to the Twitch chat channel (To see the messages! BibleThump)
5. Run `./trofiebot run` from within **$GOPATH/bin**
  - Why? Needs to access the `twitch_pass.txt` 

# Based On

Thank you to those who have helped us create this bot:

- [Base Go IRC Bot](https://github.com/Vaultpls/Twitch-IRC-Bot) >> Thanks [Vaultpls](https://github.com/Vaultpls)
- [Twitch IRC API](https://github.com/justintv/Twitch-API/blob/master/IRC.md)
- [Twitch API](https://github.com/justintv/twitch-api)
