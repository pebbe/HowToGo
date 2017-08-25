// This is a generated file. Do not edit.

m4_changequote([[, ]])
m4_define([[m4err]],[[err = $*
  if err != nil {
    fmt.Fprintf(os.Stderr, "m4_builtin([[__file__]]):m4_builtin([[__line__]]): error: %v\n", err)
    return
  }]])
