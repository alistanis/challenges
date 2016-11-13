# Lightboard

The goal of this problem is as follows:

Given an array of Toggles that control lights on a switchboard, with every combination turning on at least one light (they may only be on or off):

```
    x = on, _ = off
    No switches up
    -----------------
    | _ | _ | _ | _ |
    -----------------
    
    LightBoard - 0 = active light
    index 0 is on
    -----------------
    | x | _ | _ | _ |
    -----------------
    | _ | _ | _ | _ |
    -----------------
    | _ | _ | _ | _ |
    -----------------
    | _ | _ | _ | _ |
    -----------------
    
    Single Switch up in first position
    -----------------
    | x | _ | _ | _ |
    -----------------
    
    LightBoard
    index 1 is on
    -----------------
    | _ | x | _ | _ |
    -----------------
    | _ | _ | _ | _ |
    -----------------
    | _ | _ | _ | _ |
    -----------------
    | _ | _ | _ | _ |
    -----------------    
    
```

Given an interface Toggle that has three functions:

```
    type Toggler interface {
    	// Flips a switch in the given integer position
    	Toggle(int)
    	// Returns the current light index that is on, or -1 if no lights are on
    	Check() int
    	// The number of Toggles
    	NumToggles() int
    }
```

Write a function to test that every combination of toggles turns on only one light.