# GoWard

#### A robust and rapidly-deployable Red Team proxy with strong OPSEC considerations.

## __Description__

Generally, Red Teams and adversary's redirect their traffic through proxies to protect their backend infrastructure. GoWard proxies HTTP C2 traffic to specified Red Team servers based on the HTTP header of the traffic.

GoWard's intent is to help obfuscate Red Team traffic and provide some level of resiliency against Blue Team investigation and mitigation.

```

                   _        _
                  | \__/\__/ |
           ___    |  '.||.'  |             _
          / _ \___|__/ || \__|__ _ _ __ __| |
         / /_\/ _ \--\ || /--/ _' | '__/ _' |
        / /_\\ (_) \  \||/  / (_| | | | (_| |
        \____/\___/ \  ||  / \__,_|_|  \__,_|
                     '.||.'


                        GoWard (v0.0.1)


Usage of GoWard.exe:
  -password string
        Required.
        Specify the password for the admin panel.
                Ex: -password=pass
  -proxies int
        Required.
        Specify the number of proxies.
                Ex: -proxies=3
  -target string
        Optional.
        Specify a target URL to impersonate and use. If none specificed, default will be used.
                Ex: -target=https://www.somewebsite.co/
```

## __Features__

- Dynamically proxies traffic based on HTTP header.

- Portable and rapidly-deployable.

- Serves an impersonated webpage on port 80, which can be randomly selected from a list of real websites or specified by the user upon startup.

- Logs web requests, admin panel access, and admin panel login attempts.

- Obfuscated admin panel for alternate means of remote administration through website.

- Periodic health checks with backend infrastructure.

## __Basic Usage__

### Getting Started

GoWard is compatible with both Windows and Linux (Thanks to Go), just specify the host OS when it's built.

Once compiled and on the host which will serve as the proxy, start the program with the desired configurations. The "password" and the "proxies" fields are required.

Currently, the options are:

- __-password:__ The login password for the admin panel. Non-persistent and no default password.

- __-proxies:__ The number of proxies to be configured. After startup, user will input proxy information one at a time.

- __-target:__ The URL of a real website to impersonate. _NOTE_: Understand the potential implications of impersonating another webpage and ensure proper permissions have been received before doing so.

```
$ GoWard -password=P@ssword1 -proxies=3 -target=https://www.somewebsite.co
```

### Expected Output

Upon startup, GoWard will accept the proxy information from the user before starting the server:

```
[...]
Enter header for proxy 1: notsuspicious.com
Enter IP followed by port for proxy 1 (I.E. http://IP:PORT): http://192.168.1.244:9001
[...]
```

Additionally, GoWard will generate a log file. Verbose program output can be found there. In the console, GoWard will display what site it impersonated, along with periodic backend health checks.

```
[...]
Server started. For more verbose output, see log file: 20211231_GoWard.log
Serving impersonated webpage: https://www.somewebsite.co
Admininstration panel can be remotely accessed at /LbuBIxg/GlHglfShxH/WuWvib/tKzVlx
[...]
```

### Webpage Impersonation

GoWard will either use the provided URL to impersonate the webpage or, if none provided, randomly select a URL from the string array in `server/init.go`.

By serving an actual webpage, GoWard can help improve a Red Team's OPSEC by providing better resiliency against investigation.

### Administration Panel

GoWard will randomly generate an administration panel link upon every start-up (not persistent between sessions). Navigating to this link will provide an alternate means to remotely administrate the proxy.

Currently, the admin panel supports:

- Total web requests

- Backend infrastructure status

### C2 Framework Compatibility

GoWard has been tested successfully with:

- [Metasploit](https://github.com/rapid7/metasploit-framework)

- [Mythic](https://github.com/its-a-feature/Mythic)

_Note: Although it may not be listed, GoWard should function properly with any C2 Framework that allows for alteration of implant HTTP headers._

## __Future Features__

- Admin panel upgrades:

   - Display current proxied connections

   - More verbose data comp on web requests and connection statuses.

   - _Long Term_: Quick-reference status tracking for implants.

   - _Long Term_: Alter HTTP header and proxy information on the fly.

- HTTPS C2 traffic support

- Flag for .txt file input of header/proxy information

- Flag for .txt file input of target URLs to impersonate

- _Long Term_: CLI implementation

- More!

### __Versions__

__0.0.1:__

- Initial release

### __References__

- [Black Hat Go](https://nostarch.com/blackhatgo) - Useful information and examples for offensive security usages of Go.

- [Gorilla Web Toolkit](https://github.com/gorilla) - A lot of useful Go modules for HTTP-based applications.

### __Disclaimer__

This open source project is meant to be used with explicit authorization from any entity it affects (perceived or actual). This programs use in conjunction with offensive security tools should only take place in an approved assessment of an organization's security or for authorized research. Misuse of this software is not the responsibility of the author.
