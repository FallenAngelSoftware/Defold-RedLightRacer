components {
  id: "ScreenFaded"
  component: "/Screen-Loader/ScreenFade.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "ScreenFade"
  type: "sprite"
  data: "tile_set: \"/data/atlases/TitleImages.atlas\"\n"
  "default_animation: \"ScreenFade\"\n"
  "material: \"/Screen-Loader/test.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 320.0
    z: 1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
