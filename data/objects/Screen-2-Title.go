components {
  id: "Title-Screen"
  component: "/Screen-2-Title/Screen2.script"
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
  id: "Version"
  component: "/data/labels/Version.label"
  position {
    x: 180.0
    y: 627.0
    z: 0.01
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "Copyright"
  component: "/data/labels/Copyright.label"
  position {
    x: 180.0
    y: 16.0
    z: 0.01
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "ButtonSTARTText"
  component: "/data/labels/ButtonSTARTText.label"
  position {
    x: 180.0
    y: 450.0
    z: 0.05
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "ButtonOptionsText"
  component: "/data/labels/ButtonOptionsText.label"
  position {
    x: 180.0
    y: 400.0
    z: 0.05
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "ButtonHowToPlayText"
  component: "/data/labels/ButtonHowToPlayText.label"
  position {
    x: 180.0
    y: 350.0
    z: 0.05
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "ButtonHighScoresText"
  component: "/data/labels/ButtonHighScoresText.label"
  position {
    x: 180.0
    y: 300.0
    z: 0.05
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "ButtonAboutText"
  component: "/data/labels/ButtonAboutText.label"
  position {
    x: 180.0
    y: 250.0
    z: 0.05
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "ButtonExitText"
  component: "/data/labels/ButtonExitText.label"
  position {
    x: 180.0
    y: 200.0
    z: 0.05
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "RLR-Logo"
  type: "sprite"
  data: "tile_set: \"/data/atlases/RLR-Logo.atlas\"\n"
  "default_animation: \"RLR-Logo\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 558.0
    z: 0.01
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "Title-BG"
  type: "sprite"
  data: "tile_set: \"/data/atlases/Title-BG.atlas\"\n"
  "default_animation: \"Title-BG\"\n"
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
  id: "ButtonExit"
  type: "sprite"
  data: "tile_set: \"/data/atlases/Button.atlas\"\n"
  "default_animation: \"Button\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 200.0
    z: 0.01
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "ButtonAbout"
  type: "sprite"
  data: "tile_set: \"/data/atlases/Button.atlas\"\n"
  "default_animation: \"Button\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 250.0
    z: 0.01
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "ButtonHighScores"
  type: "sprite"
  data: "tile_set: \"/data/atlases/Button.atlas\"\n"
  "default_animation: \"Button\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 300.0
    z: 0.01
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "ButtonHowToPlay"
  type: "sprite"
  data: "tile_set: \"/data/atlases/Button.atlas\"\n"
  "default_animation: \"Button\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 350.0
    z: 0.01
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "ButtonOptions"
  type: "sprite"
  data: "tile_set: \"/data/atlases/Button.atlas\"\n"
  "default_animation: \"Button\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 400.0
    z: 0.01
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "ButtonSTART"
  type: "sprite"
  data: "tile_set: \"/data/atlases/Button.atlas\"\n"
  "default_animation: \"Button\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 450.0
    z: 0.01
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "GetItOnGooglePlay"
  type: "sprite"
  data: "tile_set: \"/data/atlases/GetItOnGooglePlay.atlas\"\n"
  "default_animation: \"GooglePlayLogo\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 180.0
    y: 90.0
    z: 0.5
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
