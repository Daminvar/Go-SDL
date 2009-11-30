
package ttf

// #include <SDL/SDL_ttf.h>
import "C";
import "sdl";
import "unsafe";


type Font struct {
    cfont *C.TTF_Font;
}

func Init() int {
    return int(C.TTF_Init());
}

func WasInit() int {
    if C.TTF_WasInit() == 1 {
        return 1
    }
    return 0;
}

func Quit() {
    C.TTF_Quit();
}

func OpenFont(file string, ptsize int) *Font {
    cfile := C.CString(file);
    cfont := C.TTF_OpenFont(cfile, C.int(ptsize));
    C.free(unsafe.Pointer(cfile));
    font := &Font{cfont};
    return font;
}

func OpenFontIndex(file string, ptsize int, index int) *Font {
    cfile := C.CString(file);
    cfont := C.TTF_OpenFontIndex(cfile, C.int(ptsize), C.long(index));
    C.free(unsafe.Pointer(cfile));
    font := &Font{cfont};
    return font;
}

func (f *Font) Close() {
    C.TTF_CloseFont(f.cfont);
}

func RenderText_Solid(font *Font, text string, color sdl.Color) *sdl.Surface {
    ctext := C.CString(text);
    ccol := C.SDL_Color{C.Uint8(color.R),C.Uint8(color.G),C.Uint8(color.B), C.Uint8(color.Unused)};
    surface := C.TTF_RenderText_Solid(font.cfont, ctext, ccol);
    C.free(unsafe.Pointer(ctext));
    return (*sdl.Surface)(unsafe.Pointer(surface));
}

func RenderText_Shaded(font *Font, text string, color sdl.Color, bgcolor sdl.Color) *sdl.Surface {
    ctext := C.CString(text);
    ccol := C.SDL_Color{C.Uint8(color.R),C.Uint8(color.G),C.Uint8(color.B), C.Uint8(color.Unused)};
    cbgcol := C.SDL_Color{C.Uint8(bgcolor.R),C.Uint8(bgcolor.G),C.Uint8(bgcolor.B), C.Uint8(bgcolor.Unused)};
    surface := C.TTF_RenderText_Shaded(font.cfont, ctext, ccol, cbgcol);
    C.free(unsafe.Pointer(ctext));
    return (*sdl.Surface)(unsafe.Pointer(surface));
}

func RenderText_Blended(font *Font, text string, color sdl.Color) *sdl.Surface {
    ctext := C.CString(text);
    ccol := C.SDL_Color{C.Uint8(color.R),C.Uint8(color.G),C.Uint8(color.B), C.Uint8(color.Unused)};
    surface := C.TTF_RenderText_Blended(font.cfont, ctext, ccol);
    C.free(unsafe.Pointer(ctext));
    return (*sdl.Surface)(unsafe.Pointer(surface));
}

func (f *Font) GetFontStyle() int {
    return int(C.TTF_GetFontStyle(f.cfont));
}

func (f *Font) SetFontStyle(style int) {
    C.TTF_SetFontStyle(f.cfont, C.int(style));
}

func (f *Font) FontHeight() int {
	return int(C.TTF_FontHeight(f.cfont));
}

func (f *Font) FontAscent() int {
	return int(C.TTF_FontAscent(f.cfont));
}

func (f *Font) FontDescent() int {
	return int(C.TTF_FontDescent(f.cfont));
}

func (f *Font) FontLineSkip() int {
	return int(C.TTF_FontLineSkip(f.cfont));
}
