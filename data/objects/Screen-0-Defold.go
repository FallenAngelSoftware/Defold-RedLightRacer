components {
  id: "Defold"
  component: "/data/gui/Screen-0-Defold-GUI.gui"
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
  id: "Screen-0-Defold"
  component: "/Screen-0-Defold/Screen0.script"
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
    z: 0.5
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "Defold-Logo"
  type: "sprite"
  data: "tile_set: \"/data/atlases/Defold-Logo.atlas\"\n"
  "default_animation: \"Defold-Logo\"\n"
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
