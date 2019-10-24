# Bing桌面

将Bing壁纸设置为电脑桌面，目前仅支持Windows和MacOS。

## 开机自启动

### Windows
将可执行文件的快捷方式放到`C:\ProgramData\Microsoft\Windows\Start Menu\Programs\StartUp`目录下，即可实现开机自启动。

### MacOs

## 定时任务

### Windows

### MacOS
推荐使用crontab进行任务管理。

```
# 每天12点00分执行
00 12 * * * /path/to/BingWallpaper

# 每月1号的12点00分执行
00 12 1 * * /path/to/BingWallpaper

# 每周一的12点00分执行
00 12 * * 1 /path/to/BingWallpaper
```
