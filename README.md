# mirror-apt
lazy apt package mirror downloader

docker build -t registry.lisong.pub:5000/sunrise/apt-mirror:1.0.3 .
docker push registry.lisong.pub:5000/sunrise/apt-mirror:1.0.3


TARGET_PATH=/tmp
SOURCE_K8S=kubernetes,https://packages.cloud.google.com/apt
SOURCE_DOCKER=docker-ce,https://download.docker.com/linux/ubuntu
SOURCE_CEPH=ceph,https://download.ceph.com/debian-nautilus
SOURCE_UBUNTU=ubuntu,http://archive.ubuntu.com/ubuntu

"ubuntu":          "https://mirrors.aliyun.com",
"debian":          "https://mirrors.ustc.edu.cn",
"debian-security": "https://mirrors.ustc.edu.cn",
"pve":             "http://download.proxmox.com/debian",
"corosync-3":      "http://download.proxmox.com/debian",
"ceph-nautilus":   "http://download.proxmox.wiki/debian",
"docker-ce":       "https://mirrors.ustc.n",