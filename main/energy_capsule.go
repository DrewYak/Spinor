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
}
