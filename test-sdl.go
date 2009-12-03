package main

import (
	"sdl";
	"sdl/ttf";
	"sdl/mixer";
)


func main() {
	sdl.Init(sdl.INIT_EVERYTHING);
	ttf.Init();
	mixer.OpenAudio(mixer.DEFAULT_FREQUENCY, mixer.DEFAULT_FORMAT,
		mixer.DEFAULT_CHANNELS, 4096);

	var screen = sdl.SetVideoMode(640, 480, 32, 0);
	sdl.WM_SetCaption("Go-SDL SDL Test", "");

	image := sdl.Load("test.png");
	running := true;
	var x, y int16;
	font := ttf.OpenFont("Fontin Sans.otf", 72);
	font.SetFontStyle(ttf.STYLE_UNDERLINE);
	white := sdl.Color{255, 255, 255, 0};
	text := ttf.RenderText_Blended(font, "Test (with music)", white);
	music := mixer.LoadMUS("test.ogg");
	music.PlayMusic(-1);

    if(sdl.GetKeyName(270)!="[+]")
    {
        panic("GetKeyName broken");
    }

	for running {

		x++;
		y++;

		e := &sdl.Event{};

		for e.Poll() {
			switch e.Type {
			case sdl.QUIT:
				running = false;
				break;
			case sdl.KEYDOWN:
				println(e.Keyboard().Keysym.Sym,": ",sdl.GetKeyName(sdl.Key(e.Keyboard().Keysym.Sym)))
			case sdl.MOUSEBUTTONDOWN:
				println("Click:", e.MouseButton().X, e.MouseButton().Y);
				x = int16(e.MouseButton().X - 16);
				y = int16(e.MouseButton().Y - 16);
			default:
			}
		}

		screen.FillRect(nil, 0x302019);
		screen.Blit(&sdl.Rect{0, 0, 0, 0}, text, nil);
		screen.Blit(&sdl.Rect{x, y, 0, 0}, image, nil);
		screen.Flip();
		sdl.Delay(25);
	}

	image.Free();
	music.Free();
	font.Close();

	ttf.Quit();
	sdl.Quit();
}
