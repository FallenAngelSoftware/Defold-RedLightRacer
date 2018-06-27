components {
  id: "Screen1Fade"
  component: "/Screen-1-FAS/ScreenFade.script"
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
  data: "tile_set: \"/data/atlases/ScreenFade.atlas\"\n"
  "default_animation: \"ScreenFade\"\n"
  "material: \"/Screen-3-Options/test.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 320.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
