components {
  id: "Screen-5-HighScores"
  component: "/data/gui/Screen-5-HighScores.gui"
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
components {
  id: "ScreenHeader"
  component: "/data/labels/ScreenHeader.label"
  position {
    x: 180.0
    y: 625.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "Screen5"
  component: "/Screen-5-HighScores/Screen5.script"
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
components {
  id: "MouseXY"
  component: "/data/labels/MouseXY.label"
  position {
    x: 5.0
    y: 630.0
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
  id: "BG"
  type: "sprite"
  data: "tile_set: \"/data/atlases/TitleFaded-BG.atlas\"\n"
  "default_animation: \"TitleFaded-BG\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
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
embedded_components {
  id: "LineTop"
  type: "sprite"
  data: "tile_set: \"/data/atlases/ScreenLine.atlas\"\n"
  "default_animation: \"ScreenLine\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 607.0
    z: 0.25
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "LineBottom"
  type: "sprite"
  data: "tile_set: \"/data/atlases/ScreenLine.atlas\"\n"
  "default_animation: \"ScreenLine\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 55.0
    z: 0.25
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
