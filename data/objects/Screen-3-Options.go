components {
  id: "Screen-3-Options"
  component: "/data/gui/Screen-3-Options.gui"
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
  id: "Screen3"
  component: "/Screen-3-Options/Screen3.script"
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
components {
  id: "OptionsMusicVolume"
  component: "/data/labels/OptionsMusicVolume.label"
  position {
    x: 305.0
    y: 560.0
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
  id: "OptionsEffectsVolume"
  component: "/data/labels/OptionsEffectsVolume.label"
  position {
    x: 305.0
    y: 510.0
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
  id: "OptionsGameMode"
  component: "/data/labels/OptionsGameMode.label"
  position {
    x: 305.0
    y: 430.0
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
  id: "OptionsMusicVolumeText"
  component: "/data/labels/OptionsMusicVolumeText.label"
  position {
    x: 55.0
    y: 560.0
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
  id: "OptionsEffectsVolumeText"
  component: "/data/labels/OptionsEffectsVolumeText.label"
  position {
    x: 55.0
    y: 510.0
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
  id: "OptionsGameModeText"
  component: "/data/labels/OptionsGameModeText.label"
  position {
    x: 55.0
    y: 430.0
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
embedded_components {
  id: "LineMidde1"
  type: "sprite"
  data: "tile_set: \"/data/atlases/ScreenLine.atlas\"\n"
  "default_animation: \"ScreenLine\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 470.0
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
  id: "LineMidde2"
  type: "sprite"
  data: "tile_set: \"/data/atlases/ScreenLine.atlas\"\n"
  "default_animation: \"ScreenLine\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 390.0
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
  id: "MenuMove"
  type: "sound"
  data: "sound: \"/data/audio/MenuMove.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
  ""
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
  id: "MenuClick"
  type: "sound"
  data: "sound: \"/data/audio/MenuClick.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
  ""
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
