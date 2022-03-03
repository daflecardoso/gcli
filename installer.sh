cd ~/
DIR="gcli"
if [ -d "$DIR" ]; then
  sudo rm -R .gcli
  echo "😳 Removing dir ${DIR}..."
fi
echo "😜 Clone project"
git clone https://github.com/daflecardoso/gcli.git .gcli
echo "🤩 Giving permission"
sudo chmod 777 .gcli
cd .gcli
echo "🧐 Installing dependencies"
npm install
npm install -g
echo "😊 Finished"
gcli
