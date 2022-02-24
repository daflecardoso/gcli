cd ~/
DIR="gcli"
if [ -d "$DIR" ]; then
  sudo rm -R gcli
  echo "Removing dir ${DIR}..."
fi

git clone https://github.com/daflecardoso/gcli.git
cd gcli
npm install -g
gcli
