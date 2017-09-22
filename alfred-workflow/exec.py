# encoding: utf-8

import subprocess
import sys
from workflow import Workflow, ICON_WEB, web


def main(wf):
    password_length = 20
    use_special = False

    if len(wf.args):
        for arg in wf.args[0].split():
            if arg.isdigit():
                password_length = int(arg)
            elif arg == 's':
                use_special = True

    if password_length < 5 or password_length > 256:
        wf.add_item(title="Password length out of bounds",
                    subtitle="Please use a reasonable password length between 5 and 256")
        wf.send_feedback()
        return 1

    command = ["./password_darwin_amd64", "get", "-l", str(password_length)]
    if use_special:
        command.append("-s")
    result = subprocess.check_output(command).strip()

    wf.add_item(title=result, arg=result)
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
