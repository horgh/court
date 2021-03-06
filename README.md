This is a Slack bot. It is primarily for demonstration purposes.

The bot listens for HTTP requests from Slack's [Events
API](https://api.slack.com/events-api). When it receives an event, it can
take action using Slack's [Web
API](https://api.slack.com/bot-users#methods) methods.


# Programs

There are two programs:

1. yorick (cmd/yorick): The Slack bot.
   * When run, it starts an HTTP server and listens for Events API HTTP
     requests.
2. horatio (cmd/horatio): An IRC bot that acts as both a Slack Events API
   and a Slack Web API.
   * It connects to an IRC server and joins a channel. It sends Events
     API-like HTTP requests for each message in the channel. It also runs
     an HTTP server where it listens for Web API-like HTTP requests to send
     messages to the channel.
   * Why? This is so we do not have to depend on configuring a bot in
     Slack's interface, nor having a Slack workspace accessible. This is
     useful for getting started quickly and in the event of workspace
     issues.


# Supported Events API events

Currently the bot knows about two events:

1. [url_verification](https://api.slack.com/events/url_verification)
   (required to configure the bot in Slack's API)
2. [message](https://api.slack.com/events/message) (a channel message)


# Supported Web API methods

Its only action is to post a message in a channel using the
[chat.postMessage](https://api.slack.com/methods/chat.postMessage) method.


# Adding your bot to a Slack workspace

You first need to get yorick running somewhere Slack will be able to send
HTTP requests.

1. Go to [api.slack.com/slack-apps](https://api.slack.com/slack-apps) and
   choose Create a Slack app
2. Pick an App Name and a Development Slack Workspace
3. Go to Event Subscriptions under Features in the left hand menu
4. Click Add a bot user and do that
5. Go to Event Subscriptions again
6. Enter your Request URL (this is yorick's endpoint URL)
7. Under Subscribe to Bot Events, choose Add Bot User Event
8. Choose message.channels
9. Go to Install App under Settings in the left hand menu
10. Choose Install App to Workspace and authorize it
11. Run yorick with the token listed as Bot User OAuth Access Token
12. In Slack, invite the bot to a channel (/invite @bot_name)
13. You should see your bot join and you can now interact with it
