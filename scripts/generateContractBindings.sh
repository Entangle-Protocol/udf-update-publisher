#!/bin/bash

# This script generates Ethereum-compatible golang bindigns for the smart contracts

usage() {
	echo "Usage: $0 --inputdir|-i <inputdir> [--pkg|-p <pkgname>] [--out|-o <outdir>]
  Generates contract bindings for each *.abi file in the provided directory

  Arguments:
    * --inputdir|-i: Directory containing *.abi files
    * --pkg|-p: Package name for the generated bindings (default: contracts)
    * --out|-o: Output directory for the generated bindings (default: contrib/contracts)
"
	exit 1
}

ABIGEN=abigen
inputdir=""
pkgname="contracts"
outdir="contrib/contracts"

# Parse arguments
while [[ $# -gt 0 ]]; do
	case "$1" in
	--inputdir | -i)
		if [[ $# -lt 2 ]]; then
			echo "Error: Missing argument for $1"
			usage
		fi
		inputdir="$2"
		shift 2
		;;
	--pkg | -p)
		if [[ $# -lt 2 ]]; then
			echo "Error: Missing argument for $1"
			usage
		fi
		pkgname="$2"
		shift 2
		;;
	--out | -o)
		if [[ $# -lt 2 ]]; then
			echo "Error: Missing argument for $1"
			usage
		fi
		outdir="$2"
		shift 2
		;;
	*)
		echo "Error: Invalid argument: $1"
		usage
		;;
	esac
done

if [[ -z "$inputdir" ]]; then
  echo "Error: Missing input directory"
  usage
fi

echo "pkgname: $pkgname"
echo "outdir: $outdir"
echo "inputdir: $inputdir"

abiFiles=$(find "$inputdir" -name "*.abi")

for abiFile in $abiFiles; do
  echo "Processing $abiFile"
  abiName="$(basename "$abiFile" .abi)"

  # Parse providerName from directory name
  dirName="$(dirname "$abiFile")"
  providerName=${dirName##*abi/}

  # Create the output directory if it doesn't exist
  outProviderPath="$outdir/$providerName/$abiName"
  mkdir -p "$outProviderPath"

  # Generate golang contract bindings for the current abi file
  $ABIGEN --abi "$abiFile" --pkg "$abiName" --out "$outProviderPath/$abiName.go"
done
