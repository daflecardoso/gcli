cd ~/
DIR="gcli"
if [ -d "$DIR" ]; then
  sudo rm -R gcli
  echo "Removing dir ${DIR}..."
fi

mkdir gcli
sudo chmod 777 gcli
cd gcli
git clone https://github.com/daflecardoso/gcli.git
cd gcli
npm install
gcli
