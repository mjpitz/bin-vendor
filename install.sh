readonly version="${VERSION}"
os=$(uname -s | tr '[:upper:]' '[:lower:]')
arch=$(uname -m)

if [[ -z "${version}" ]]; then
  echo "must set VERSION"
  exit 1
fi

if [[ "x86_64" == "${arch}" ]]; then
  arch="amd64"
elif [[ "aarch64" == "${arch}" ]]; then
  arch="arm64"
else
  echo "unsupported machine architecture"
  exit 1
fi

readonly asset_name="bin-vendor_${version##v}_${os}_${arch}.tar.gz"
readonly asset_url="https://github.com/mjpitz/bin-vendor/releases/download/${version}/${asset_name}"

pushd "$(mktemp -d)" || exit 1
curl -sSLo bin-vendor.tar.gz "${asset_url}"
tar zxf bin-vendor.tar.gz
sudo install bin-vendor /usr/local/bin
popd || exit 1
