components {
  id: "damageHero"
  component: "/main/hero/damageHeroText.script"
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
  id: "heroLoseHealth"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "  z: 0.0\n"
  "  w: 0.0\n"
  "}\n"
  "scale {\n"
  "  x: 0.572\n"
  "  y: 0.572\n"
  "  z: 0.572\n"
  "  w: 0.0\n"
  "}\n"
  "color {\n"
  "  x: 0.9019608\n"
  "  y: 0.3019608\n"
  "  z: 0.3019608\n"
  "  w: 1.0\n"
  "}\n"
  "outline {\n"
  "  x: 0.6\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "shadow {\n"
  "  x: 0.4\n"
  "  y: 0.4\n"
  "  z: 0.4\n"
  "  w: 1.0\n"
  "}\n"
  "leading: 1.0\n"
  "tracking: 0.0\n"
  "pivot: PIVOT_CENTER\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "line_break: false\n"
  "text: \"heroLoseHealth\"\n"
  "font: \"/builtins/fonts/system_font.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
