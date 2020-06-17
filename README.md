# go-rtaudio

RTAudio bindings for go - with asio/wasapi support on Windows   
Original bindings from the rtaudio source in contrib


# Compile RTAudio on Windows with asio/wasapi

### in the source dir make a build directory
> mkdir _buildmsys   
### make
> cmake -G 'MSYS Makefiles' -DRTAUDIO_API_ASIO=ON -DRTAUDIO_API_DS=ON ..   
### or   
> cmake -G 'MinGW Makefiles' -DRTAUDIO_API_ASIO=ON -DRTAUDIO_API_DS=ON ..  


## build   
> mingw32-make.exe   
> mingw32-make.exe install prefix=/mingw64


## check your rtaudio.pc file in /mingw64/lib/pkg-config
It should look like this :   
``` 
prefix=/mingw64
exec_prefix=${prefix}
libdir=${exec_prefix}/lib
includedir=${prefix}/include/rtaudio        

Name: librtaudio
Description: RtAudio - a set of C++ classes that provide a common API for realtime audio input/output
Version: 5.1.0
Requires:  
Libs: -L${libdir} -lrtaudio -lm -luuid -lksuser -lwinmm -lole32 -lcomdlg32
Libs.private: -lpthread
Cflags: -pthread -I${includedir}  -D__WINDOWS_ASIO__ -D__WINDOWS_WASAPI__ 
```
