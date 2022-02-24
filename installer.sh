cd ~/
DIR="gcli"
if [ -d "$DIR" ]; then
  sudo rm -R gcli
  echo "Removing dir ${DIR}..."
fi

git clone https://github.com/daflecardoso/gcli.git
sudo chmod 777 -R
cd gcli
npm install -g
gcli
