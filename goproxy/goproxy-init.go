package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	proxyHost       = "localhost:8080" // 代理服务器地址
	moduleCachePath = "./module_cache" // 模块缓存目录
)

func main() {
	// 创建模块缓存目录
	if err := os.MkdirAll(moduleCachePath, os.ModePerm); err != nil {
		log.Fatal("Failed to create module cache directory:", err)
	}

	// 设置 HTTP 请求处理函数
	http.HandleFunc("/", proxyHandler)

	// 启动代理服务器
	fmt.Println("Go Proxy server is running on", proxyHost)
	log.Fatal(http.ListenAndServe(proxyHost, nil))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// 获取模块信息
	moduleName := r.URL.Path[1:]

	// 拼接本地缓存路径
	cachePath := filepath.Join(moduleCachePath, moduleName)

	// 检查本地是否有缓存文件
	if _, err := os.Stat(cachePath); err == nil {
		// 缓存存在，直接返回缓存文件
		fmt.Println("Serving module from cache:", moduleName)
		serveFile(w, r, cachePath)
		return
	}

	// 缓存不存在，从远程模块源下载模块文件
	fmt.Println("Downloading module:", moduleName)
	downloadModule(moduleName, cachePath)

	// 返回下载的模块文件给客户端
	serveFile(w, r, cachePath)
}

func downloadModule(moduleName string, outputPath string) {
	// 构造远程模块源的请求URL
	moduleURL := "https://proxy.golang.org/" + moduleName

	// 发起 GET 请求，获取模块文件
	resp, err := http.Get(moduleURL)
	if err != nil {
		log.Println("Failed to download module:", err)
		return
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		log.Println("Failed to download module. Status code:", resp.StatusCode)
		return
	}

	// 创建本地模块文件
	file, err := os.Create(outputPath)
	if err != nil {
		log.Println("Failed to create module file:", err)
		return
	}
	defer file.Close()

	// 将模块文件保存到本地
	if _, err := io.Copy(file, resp.Body); err != nil {
		log.Println("Failed to save module file:", err)
		return
	}

	fmt.Println("Module downloaded:", moduleName)
}

func serveFile(w http.ResponseWriter, r *http.Request, filePath string) {
	// 检查文件是否存在
	_, err := os.Stat(filePath)
	if err != nil {
		http.NotFound(w, r) // 文件不存在，返回 404 Not Found
		return
	}

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Could not get file info", http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", string(fileInfo.Size()))

	// 读取文件内容并写入响应体
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Could not write file content", http.StatusInternalServerError)
		return
	}
}