package sdl

// struct private_hwdata{};
// struct SDL_BlitMap{};
// #define map _map
// #include <SDL/SDL.h>
// #include <SDL/SDL_image.h>
// #include <SDL/SDL_ttf.h>
import "C"
import "unsafe"

import "image"

//type Surface C.SDL_Surface;
//type TTF_Font C.TTF_Font;

type cast unsafe.Pointer

//SDL

func Init(flags uint32) int	{ return int(C.SDL_Init(C.Uint32(flags))) }
func Quit()	{ C.SDL_Quit() }

//SDL_video

func SetVideoMode(w int, h int, bpp int, flags uint32) *Surface {
	var screen = C.SDL_SetVideoMode(C.int(w), C.int(h), C.int(bpp), C.Uint32(flags));
	return (*Surface)(cast(screen));
}

func WM_SetCaption(title string, icon string) {
    ctitle:=C.CString(title);
    cicon:=C.CString(icon);
    C.SDL_WM_SetCaption(ctitle,cicon);
    C.free(unsafe.Pointer(ctitle));
    C.free(unsafe.Pointer(cicon));
}

func GL_SwapBuffers()
{
    C.SDL_GL_SwapBuffers();
}


func (screen *Surface) Flip() int	{ return int(C.SDL_Flip((*C.SDL_Surface)(cast(screen)))) }

func (screen *Surface) Free()	{ C.SDL_FreeSurface((*C.SDL_Surface)(cast(screen))) }

func (screen *Surface) Lock() int {
	return int(C.SDL_LockSurface((*C.SDL_Surface)(cast(screen))))
}

func (screen *Surface) Unlock() int	{ return int(C.SDL_Flip((*C.SDL_Surface)(cast(screen)))) }

func GetVideoSurface() *Surface {
	var screen = C.SDL_GetVideoSurface();
	return (*Surface)(cast(screen));
}

func (dst *Surface) Blit(dstrect *Rect, src *Surface, srcrect *Rect) int {
	var ret = C.SDL_UpperBlit(
		(*C.SDL_Surface)(cast(src)),
		(*C.SDL_Rect)(cast(srcrect)),
		(*C.SDL_Surface)(cast(dst)),
		(*C.SDL_Rect)(cast(dstrect)));

	return int(ret);
}

func (dst *Surface) FillRect(dstrect *Rect, color uint32) int {
	var ret = C.SDL_FillRect(
		(*C.SDL_Surface)(cast(dst)),
		(*C.SDL_Rect)(cast(dstrect)),
		C.Uint32(color));

	return int(ret);
}

// surface --> Image

func (surface *Surface) ColorModel() image.ColorModel
{
    return nil;
}

func (surface *Surface) Width() int
{
    return int(surface.W);
}

func (surface *Surface) Height() int
{
    return int(surface.H);
}

func (surface *Surface) At(x, y int) image.Color
{

    //TODO endianess, bpp, alpha, etc

    var bpp = int(surface.Format.BytesPerPixel);

    var pixel=uintptr(unsafe.Pointer(surface.Pixels));

    pixel += uintptr( y * int(surface.Pitch) + x * bpp);

    var color= *((*image.RGBAColor)(unsafe.Pointer(pixel)));

    return image.RGBAColor{color.B,color.G,color.R,255-color.A};

}

//SDL image

func Load(file string) *Surface {
    cfile := C.CString(file);
	var screen = C.IMG_Load(cfile);
    C.free(unsafe.Pointer(cfile));
	return (*Surface)(cast(screen));
}

//SDL events

func (event *Event) Wait() bool {
	var ret = C.SDL_WaitEvent((*C.SDL_Event)(cast(event)));
	return ret != 0;
}

func (event *Event) Poll() bool {
	var ret = C.SDL_PollEvent((*C.SDL_Event)(cast(event)));
	return ret != 0;
}

func (event *Event) Keyboard() *KeyboardEvent
{
    if(event.Type==KEYUP || event.Type==KEYDOWN)
    {
        return (*KeyboardEvent)(cast(event));
    }

    return nil;
}

func (event *Event) MouseButton() *MouseButtonEvent
{
    if(event.Type==MOUSEBUTTONDOWN || event.Type==MOUSEBUTTONUP)
    {
        return (*MouseButtonEvent)(cast(event));
    }

    return nil;
}

func (event *Event) MouseMotion() *MouseMotionEvent
{
    if(event.Type==MOUSEMOTION)
    {
        return (*MouseMotionEvent)(cast(event));
    }

    return nil;
}

//SDL time

func Delay(ms uint32)	{ C.SDL_Delay(C.Uint32(ms)) }

//SDL ttf


func TTF_Init() int {
    return int(C.TTF_Init());
}

func TTF_Quit() {
    C.TTF_Quit();
}

func TTF_OpenFont(file string, ptsize int) *C.TTF_Font {
    cfile := C.CString(file);
    cfont := C.TTF_OpenFont(cfile, C.int(ptsize));
    C.free(unsafe.Pointer(cfile));
    return cfont;
}

func TTF_CloseFont(font *C.TTF_Font) {
    C.TTF_CloseFont(font);
}

func TTF_RenderText_Blended(cfont *C.TTF_Font, text string, color Color) *Surface {
    ctext := C.CString(text);
    ccol := C.SDL_Color{C.Uint8(color.R),C.Uint8(color.G),C.Uint8(color.B), C.Uint8(color.Unused)};
    surface := C.TTF_RenderText_Blended(cfont, ctext, ccol);
    C.free(unsafe.Pointer(ctext));
    return (*Surface)(cast(surface));
}


