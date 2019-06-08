#!python3

from sys import argv
from argparse import ArgumentParser
from subprocess import call
from subprocess import run


def update():
    call(['git', 'pull'])


def start():
    run(['screen', 'go', 'run', 'dbcd.go'])


def kill():
    call(['killall', 'dbcd'])


def restart():
    kill()
    start()


def main():
    parser = ArgumentParser(description='The script controls the status of D̲atab̲ase C̲ourse D̲esign server.')
    
    group = parser.add_mutually_exclusive_group(required=True)
    group.add_argument('-u', '--update', action='store_true', help='update code from git repo')
    group.add_argument('-s', '--start', action='store_true', help='start the server')
    group.add_argument('-k', '--kill', '--stop', action='store_true', help='kill the server')
    group.add_argument('-r', '--restart', action='store_true', help='restart the server')
    
    # Provide help if no argument is specified.
    if len(argv) < 2:
        parser.print_help()
        return

    # Error when more than one argument is passed.
    if len(argv) > 2:
        parser.error('passing more than one argument is not allowed')
        
    args = parser.parse_args()

    if args.update:
        update()
        return

    if args.start:
        start()
        return

    if args.kill:
        kill()
        return

    if args.restart:
        restart()
        return


if __name__ == "__main__":
    main()
