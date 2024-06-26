package evaluator

import (
	"fmt"
	"interpreter/object"
)

var builtins = map[string]*object.Builtin {
	"out": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
	"tail": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `tail` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}
			return NULL
		},
	},
	"head": &object.Builtin{
    Fn: func(args ...object.Object) object.Object {
      if len(args) != 1 {
          return newError("wrong number of arguments. got=%d, want=1", len(args))
      }
      if args[0].Type() != object.ARRAY_OBJ {
          return newError("argument to `init` must be ARRAY, got %s", args[0].Type())
      }
      arr := args[0].(*object.Array)
      length := len(arr.Elements)
      if length > 0 {
          newElements := make([]object.Object, length-1, length-1)
          copy(newElements, arr.Elements[:length-1])
          return &object.Array{Elements: newElements}
      }
      return NULL
    },
	},
	"prepend": &object.Builtin{
    Fn: func(args ...object.Object) object.Object {
      if len(args) != 2 {
          return newError("wrong number of arguments. got %d, want=2", len(args))
      }
      if args[0].Type() != object.ARRAY_OBJ {
          return newError("argument to `prepend` must be ARRAY, got %s", args[0].Type())
      }

      arr := args[0].(*object.Array)
      arr.Elements = append([]object.Object{args[1]}, arr.Elements...)
      return arr
    },
	},

	"append": &object.Builtin{
    Fn: func(args ...object.Object) object.Object {
      if len(args) != 2 {
          return newError("wrong number of arguments. got %d, want=2", len(args))
      }
      if args[0].Type() != object.ARRAY_OBJ {
          return newError("argument to `append` must be ARRAY, got %s", args[0].Type())
      }

      arr := args[0].(*object.Array)
      arr.Elements = append(arr.Elements, args[1])

      return arr
    },
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
						len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}
			return NULL
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
						len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s",
						args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"len": &object.Builtin {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
						len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value :int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s",
						args[0].Type())
			}
		},
	},
}
