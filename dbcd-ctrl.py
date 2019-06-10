#!python3

from sys import argv
from argparse import ArgumentParser
from subprocess import call


def main():
    # Init argument parser
    parser = ArgumentParser(description='The script controls the status of D̲atab̲ase C̲ourse D̲esign server.')
    group = parser.add_mutually_exclusive_group(required=True)
    group.add_argument('-u', '--update', action='store_true', help='update code from git repo')
    group.add_argument('-s', '--start', action='store_true', help='start the server')
    group.add_argument('-k', '--kill', '--stop', action='store_true', help='kill the server')
    group.add_argument('-r', '--restart', action='store_true', help='restart the server')
    
    # Print help if no specific argument.
    if len(argv) < 2:
        parser.print_help()
        return

    # Error when passing two or more arguments.
    if len(argv) > 2:
        parser.error('passing more than one argument is not allowed')
    
    # Parse argument.
    args = parser.parse_args()

    # Act according to parsed argument.
    if args.update:
        update_server()
        return

    if args.start:
        start_server()
        return

    if args.kill:
        kill_server()
        return

    if args.restart:
        restart_server()
        return


def update_server():
    call(['git', 'pull'])


def start_server():
    call(['go', 'run', 'dbcd.go'])


def kill_server():
    call(['killall', 'dbcd'])


def restart_server():
    kill_server()
    start_server()


if __name__ == "__main__":
    main()
