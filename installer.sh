cd ~/
DIR="gcli"
if [ -d "$DIR" ]; then
  sudo rm -R gcli
  echo "😳 Removing dir ${DIR}..."
fi
echo "😜 Clone project"
git clone https://github.com/daflecardoso/gcli.git
echo "🤩 Giving permission"
sudo chmod 777 -R gcli
cd gcli
echo "🧐 Installing dependencies"
npm install
npm install -g
echo "😊 Finished"
gcli
