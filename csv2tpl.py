#!/usr/bin/python3

# Parse a python string.Template file and replace all variables from
# the given CSV file. CSV file headers will be used as template variable names.

import csv, string, argparse, string, sys

# Step 1: parse command line args
parser = argparse.ArgumentParser()
parser.add_argument("tplfname", help="Template file name")
parser.add_argument("csvfname", help="CSV file name")
args = parser.parse_args()
csvfname = args.csvfname
tplfname = args.tplfname

# Setting this will silently ignore blank values in the CSV file
# If unset, the script will error and exit
# TODO: make this an optional cmd line argument
ignoremissing = False

# Step 2: Open template file
template = open(tplfname)
src = string.Template(template.read())

# Step 3: Open the csv file
csvfile = open(csvfname)
reader = csv.DictReader(csvfile)

# Step 4: For every row in CSV file, use the template and print
for row in reader:
    if ignoremissing:
        # Strip whitespace from keys and values
        d = {k.strip(): v.strip() for k, v in row.items()}
    else:
        # Strip whitespace from keys and values
        d = {k.strip(): v.strip() for k, v in row.items() if v}

    # Substitute variables in template
    try:
        print(src.substitute(d))
    except KeyError as err:
        print('ERROR: Missing value:', err, file=sys.stderr)
        print('ERROR: Available:', d, file=sys.stderr)
        sys.exit(1)
