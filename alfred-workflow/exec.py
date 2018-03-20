# encoding: utf-8

import subprocess
import sys
import json
from workflow import Workflow, ICON_WEB, web

descriptions = {
    "htpasswd_apr1": "APR1 (htpasswd)",
    "htpasswd_bcrypt": "bcrypt (htpasswd)",
    "htpasswd_sha256": "SHA256 (htpasswd)",
    "htpasswd_sha512": "SHA512 (htpasswd)",
    "password": "Password",
    "sha1": "SHA1",
    "sha256": "SHA256",
    "sha512": "SHA512",
}


def main(wf):
    password_length = 20
    use_special = False
    use_xkcd = False

    if len(wf.args):
        for arg in wf.args[0].split():
            if arg.isdigit():
                password_length = int(arg)
            elif arg == 's':
                use_special = True
            elif arg == 'x':
                use_xkcd = True

    if password_length < 4 or password_length > 256:
        wf.add_item(title="Password length out of bounds",
                    subtitle="Please use a reasonable password length between 4 and 256")
        wf.send_feedback()
        return 1

    command = ["./password_darwin_amd64",
               "get", "-j", "-l",
               str(password_length)]
    if use_special:
        command.append("-s")
    if use_xkcd:
        command.append("-x")
    result = json.loads(subprocess.check_output(command).strip())

    hashed = []
    for key, value in result.iteritems():
        hashed.append("{}: {}".format(key, value))

    wf.add_item(title=result['password'],
                subtitle="Press Cmd+C to copy",
                arg=result['password'],
                valid=True)
    wf.add_item(title="Copy hashed versions",
                subtitle="Press Cmd+C to copy",
                arg="\n".join(hashed),
                valid=True)
    wf.send_feedback()

    return 0


if __name__ == "__main__":
    wf = Workflow(update_settings={
        'github_slug': 'Luzifer/password',
    })

    if wf.update_available:
        # Download new version and tell Alfred to install it
        wf.start_update()

    sys.exit(wf.run(main))
