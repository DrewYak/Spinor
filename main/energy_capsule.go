components {
  id: "energy_capsule"
  component: "/main/energy_capsule.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"energy_capsule\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/game.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.08
    y: 0.08
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"energy\"\n"
  "mask: \"satellite\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 8.948364\n"
  "  data: 35.22254\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
