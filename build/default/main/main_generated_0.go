components {
  id: "character"
  component: "/main/hero/character.script"
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
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/main/hero/character.tilesource\"\ndefault_animation: \"idle\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ALPHA\n"
  position {
    x: 24.635
    y: 23.0
    z: 0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "music"
  type: "sound"
  data: "sound: \"/audio/music/mystery.wav\"\nlooping: 1\ngroup: \"music\"\ngain: 0.02\n"
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
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_KINEMATIC\nmass: 0.0\nfriction: 0.0\nrestitution: 0.5\ngroup: \"default\"\nmask: \"default\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: 24.096\n      y: 22.679\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 1\n  }\n  data: 12.0\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\n"
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
  id: "laserFactory"
  type: "factory"
  data: "prototype: \"/main/hero/laser.go\"\nload_dynamically: false\n"
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
  id: "laserHerofactory"
  type: "factory"
  data: "prototype: \"/main/hero/laserHero.go\"\nload_dynamically: false\n"
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
  id: "coinFactory"
  type: "factory"
  data: "prototype: \"/main/hero/coin.go\"\nload_dynamically: false\n"
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
  id: "ding"
  type: "sound"
  data: "sound: \"/audio/soundEffects/ding.wav\"\nlooping: 0\ngroup: \"ding\"\ngain: 0.35\n"
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
  id: "heroLoseHp"
  type: "label"
  data: "size {\n  x: 128.0\n  y: 32.0\n  z: 0.0\n  w: 0.0\n}\nscale {\n  x: 0.623\n  y: 0.623\n  z: 0.0\n  w: 0.0\n}\ncolor {\n  x: 0.6\n  y: 0.0\n  z: 0.0\n  w: 1.0\n}\noutline {\n  x: 1.0\n  y: 0.6\n  z: 0.6\n  w: 1.0\n}\nshadow {\n  x: 0.0\n  y: 0.0\n  z: 0.0\n  w: 1.0\n}\nleading: 1.0\ntracking: 0.0\npivot: PIVOT_CENTER\nblend_mode: BLEND_MODE_ALPHA\nline_break: false\ntext: \"\"\nfont: \"/builtins/fonts/system_font.font\"\nmaterial: \"/builtins/fonts/label.material\"\n"
  position {
    x: 24.231
    y: 36.349
    z: 0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "loseHealthFactory"
  type: "factory"
  data: "prototype: \"/main/hero/damageHeroText.go\"\nload_dynamically: false\n"
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
  id: "loseEnemyHealth"
  type: "factory"
  data: "prototype: \"/main/hero/damageEnemyText.go\"\nload_dynamically: false\n"
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
  id: "slimeDeath"
  type: "sound"
  data: "sound: \"/audio/soundEffects/slimeDeath.wav\"\nlooping: 0\ngroup: \"hitEffect\"\ngain: 1.0\n"
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
  id: "slimeHit"
  type: "sound"
  data: "sound: \"/audio/soundEffects/slimeHit.wav\"\nlooping: 0\ngroup: \"master\"\ngain: 1.0\n"
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
  id: "playerHit"
  type: "sound"
  data: "sound: \"/audio/soundEffects/playerHit.wav\"\nlooping: 0\ngroup: \"master\"\ngain: 0.2\n"
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
  id: "criticalHit"
  type: "sound"
  data: "sound: \"/audio/soundEffects/criticalHit.wav\"\nlooping: 0\ngroup: \"master\"\ngain: 0.1\n"
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
  id: "bulletFactory"
  type: "factory"
  data: "prototype: \"/main/hero/bullet.go\"\nload_dynamically: false\n"
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
  id: "bulletGoblin"
  type: "factory"
  data: "prototype: \"/main/hero/bulletGoblin.go\"\nload_dynamically: false\n"
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
  id: "swordSwing"
  type: "sound"
  data: "sound: \"/audio/soundEffects/swordSwing.wav\"\nlooping: 0\ngroup: \"master\"\ngain: 0.05\n"
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
  id: "laserSound"
  type: "sound"
  data: "sound: \"/audio/soundEffects/laserSound.wav\"\nlooping: 0\ngroup: \"master\"\ngain: 0.1\n"
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
