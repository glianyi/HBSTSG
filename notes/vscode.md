vscode识别不了struct addrinfo问题,addrinfo是posix标准,如果指定了-std=c99 gcc编译会有问题
vscode需要在c_cpp_properties defines属性中指定宏才能被编辑器识别
https://github.com/Microsoft/vscode-cpptools/issues/2025
https://zhuanlan.zhihu.com/p/602476352