// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ki

import (
	"reflect"

	"cogentcore.org/core/enums"
	"cogentcore.org/core/gti"
)

// The Ki interface provides the core functionality for a Cogent Core tree.
// Each Ki is a node in the tree and can have child nodes, and no cycles
// are allowed (i.e., each node can only appear once in the tree).
// All the usual methods are included for accessing and managing Children,
// and efficiently traversing the tree and calling functions on the nodes.
// In addition, Ki nodes can have Fields that are also Ki nodes that
// are included in all the automatic tree traversal methods -- they are
// effectively named fixed children that are automatically present.
//
// In general, the names of the children of a given node should all be unique.
// The following functions defined in ki package can be used:
// UniqueNameCheck(node) to check for unique names on node if uncertain.
// UniqueNameCheckAll(node) to check entire tree under given node.
// UniquifyNames(node) to add a suffix to name to ensure uniqueness.
// UniquifyNamesAll(node) to to uniquify all names in entire tree.
//
// Use function MoveChild to move a node between trees or within a tree --
// otherwise nodes are typically created and deleted but not moved.
//
// The Ki interface is designed to support virtual method calling in Go
// and is only intended to be implemented once, by the ki.Node type
// (as opposed to interfaces that are used for hiding multiple different
// implementations of a common concept).  Thus, all of the fields in ki.Node
// are exported (have captital names), to be accessed directly in types
// that embed and extend the ki.Node. The Ki interface has the "formal" name
// (e.g., Children) while the Node has the "nickname" (e.g., Kids).  See the
// Naming Conventions on the Cogent Core Wiki for more details.
//
// Each Node stores the Ki interface version of itself, as This() / Ths
// which enables full virtual function calling by calling the method
// on that interface instead of directly on the receiver Node itself.
// This requires proper initialization via Init method of the Ki interface.
//
// Ki nodes also support the following core functionality:
//   - UpdateStart() / UpdateEnd() to wrap around tree updating code, which then
//     automatically triggers update signals at the highest level of the
//     affected tree, resulting in efficient updating logic for arbitrary
//     nested tree modifications.
//   - ConfigChildren system for minimally updating children to fit a given
//     Name & Type template.
//   - Automatic JSON I/O of entire tree including type information.
type Ki interface {
	// InitName initializes this node to given actual object as a Ki interface
	// and sets its name. The names should be unique among children of a node.
	// This is needed for root nodes -- automatically done for other nodes
	// when they are added to the Ki tree. If the name is unspecified, it
	// defaults to the ID (kebab-case) name of the type.
	// Even though this is a method and gets the method receiver, it needs
	// an "external" version of itself passed as the first arg, from which
	// the proper Ki interface pointer will be obtained.  This is the only
	// way to get virtual functional calling to work within the Go language.
	InitName(this Ki, name ...string)

	// This returns the Ki interface that guarantees access to the Ki
	// interface in a way that always reveals the underlying type
	// (e.g., in reflect calls).  Returns nil if node is nil,
	// has been destroyed, or is improperly constructed.
	This() Ki

	// AsNode returns the *ki.Node base type for this node.
	AsKi() *Node

	// BaseType returns the base node type for all elements within this tree.
	// Used e.g., for determining what types of children can be created.
	BaseType() *gti.Type

	// Name returns the user-defined name of the object (Node.Nm),
	// for finding elements, generating paths, IO, etc.
	Name() string

	// SetName sets the name of this node.
	// Names should generally be unique across children of each node.
	// See Unique* functions to check / fix.
	// If node requires non-unique names, add a separate Label field.
	// Does NOT wrap in UpdateStart / End.
	SetName(name string)

	// KiType returns the gti Type record for this Ki node.
	// This is auto-generated by the gtigen generator for Ki types.
	KiType() *gti.Type

	// New returns a new token of this Ki node.
	// InitName _must_ still be called on this new token.
	// This is auto-generated by the gtigen generator for Ki types.
	New() Ki

	//////////////////////////////////////////////////////////////////////////
	//  Parents

	// Parent returns the parent of this Ki (Node.Par) -- Ki has strict
	// one-parent, no-cycles structure -- see SetParent.
	Parent() Ki

	// IndexInParent returns our index within our parent object. It caches the
	// last value and uses that for an optimized search so subsequent calls
	// are typically quite fast. Returns -1 if we don't have a parent.
	IndexInParent() int

	// ParentLevel finds a given potential parent node recursively up the
	// hierarchy, returning level above current node that the parent was
	// found, and -1 if not found.
	ParentLevel(par Ki) int

	// ParentByName finds first parent recursively up hierarchy that matches
	// given name.  Returns nil if not found.
	ParentByName(name string) Ki

	// ParentByType finds parent recursively up hierarchy, by type, and
	// returns nil if not found. If embeds is true, then it looks for any
	// type that embeds the given type at any level of anonymous embedding.
	ParentByType(t *gti.Type, embeds bool) Ki

	//////////////////////////////////////////////////////////////////////////
	//  Children

	// HasChildren tests whether this node has children (i.e., non-terminal).
	HasChildren() bool

	// NumChildren returns the number of children
	NumChildren() int

	// NumLifetimeChildren returns the number of children that this node
	// has ever had added to it (it is not decremented when a child is removed).
	// It is used for unique naming of children.
	NumLifetimeChildren() uint64

	// Children returns a pointer to the slice of children (Node.Kids) -- use
	// methods on ki.Slice for further ways to access (ByName, ByType, etc).
	// Slice can be modified, deleted directly (e.g., sort, reorder) but Add
	// method on parent node should be used to ensure proper init.
	Children() *Slice

	// Child returns the child at given index and returns nil if
	// the index is out of range.
	Child(idx int) Ki

	// ChildByName returns the first element that has given name, and nil
	// if no such element is found. startIdx arg allows for optimized
	// bidirectional find if you have an idea where it might be, which
	// can be a key speedup for large lists. If no value is specified for
	// startIdx, it starts in the middle, which is a good default.
	ChildByName(name string, startIdx ...int) Ki

	// ChildByType returns the first element that has the given type, and nil
	// if not found. If embeds is true, then it also looks for any type that
	// embeds the given type at any level of anonymous embedding.
	// startIdx arg allows for optimized bidirectional find if you have an
	// idea where it might be, which can be a key speedup for large lists. If
	// no value is specified for startIdx, it starts in the middle, which is a
	// good default.
	ChildByType(t *gti.Type, embeds bool, startIdx ...int) Ki

	//////////////////////////////////////////////////////////////////////////
	//  Paths

	// Path returns path to this node from the tree root, using node Names
	// separated by / and fields by .
	// Node names escape any existing / and . characters to \\ and \,
	// Path is only valid when child names are unique (see Unique* functions)
	Path() string

	// PathFrom returns path to this node from given parent node, using
	// node Names separated by / and fields by .
	// Node names escape any existing / and . characters to \\ and \,
	// Path is only valid for finding items when child names are unique
	// (see Unique* functions). The paths that it returns exclude the
	// name of the parent and the leading slash; for example, in the tree
	// a/b/c/d/e, the result of d.PathFrom(b) would be c/d. PathFrom
	// automatically gets the [Ki.This] version of the given parent,
	// so a base type can be passed in without manually calling [Ki.This].
	PathFrom(par Ki) string

	// FindPath returns Ki object at given path, starting from this node
	// (e.g., the root).  If this node is not the root, then the path
	// to this node is subtracted from the start of the path if present there.
	// FindPath only works correctly when names are unique.
	// Path has node Names separated by / and fields by .
	// Node names escape any existing / and . characters to \\ and \,
	// There is also support for [idx] index-based access for any given path
	// element, for cases when indexes are more useful than names.
	// Returns nil if not found.
	FindPath(path string) Ki

	// FieldByName returns Ki object that is a direct field.
	// This must be implemented for any types that have Ki fields that
	// are processed as part of the overall Ki tree.  This is only used
	// by FindPath.
	// Returns error if not found.
	FieldByName(field string) (Ki, error)

	//////////////////////////////////////////////////////////////////////////
	//  Adding, Inserting Children

	// AddChild adds given child at end of children list.
	// The kid node is assumed to not be on another tree (see MoveToParent)
	// and the existing name should be unique among children.
	// No UpdateStart / End wrapping is done: do that externally as needed.
	// Can also call SetFlag(ki.ChildAdded) if notification is needed.
	AddChild(kid Ki) error

	// NewChild creates a new child of the given type and adds it at end
	// of children list. The name should be unique among children. If the
	// name is unspecified, it defaults to the ID (kebab-case) name of the
	// type, plus the [Ki.NumLifetimeChildren] of its parent.
	// No UpdateStart / End wrapping is done: do that externally as needed.
	// Can also call SetFlag(ki.ChildAdded) if notification is needed.
	NewChild(typ *gti.Type, name ...string) Ki

	// SetChild sets child at given index to be the given item; if it is passed
	// a name, then it sets the name of the child as well; just calls Init
	// (or InitName) on the child, and SetParent. Names should be unique
	// among children. No UpdateStart / End wrapping is done: do that
	// externally as needed. Can also call SetFlag(ki.ChildAdded) if
	// notification is needed.
	SetChild(kid Ki, idx int, name ...string) error

	// InsertChild adds given child at position in children list.
	// The kid node is assumed to not be on another tree (see MoveToParent)
	// and the existing name should be unique among children.
	// No UpdateStart / End wrapping is done: do that externally as needed.
	// Can also call SetFlag(ki.ChildAdded) if notification is needed.
	InsertChild(kid Ki, at int) error

	// InsertNewChild creates a new child of given type and add at position
	// in children list. The name should be unique among children. If the
	// name is unspecified, it defaults to the ID (kebab-case) name of the
	// type, plus the [Ki.NumLifetimeChildren] of its parent. No
	// UpdateStart / End wrapping is done: do that externally as needed.
	// Can also call SetFlag(ki.ChildAdded) if notification is needed.
	InsertNewChild(typ *gti.Type, at int, name ...string) Ki

	// SetNChildren ensures that there are exactly n children, deleting any
	// extra, and creating any new ones, using NewChild with given type and
	// naming according to nameStubX where X is the index of the child.
	// If nameStub is not specified, it defaults to the ID (kebab-case)
	// name of the type.
	//
	// IMPORTANT: returns whether any modifications were made (mods) AND if
	// that is true, the result from the corresponding UpdateStart call --
	// UpdateEnd is NOT called, allowing for further subsequent updates before
	// you call UpdateEnd(updt)
	//
	// Note that this does not ensure existing children are of given type, or
	// change their names, or call UniquifyNames -- use ConfigChildren for
	// those cases -- this function is for simpler cases where a parent uses
	// this function consistently to manage children all of the same type.
	SetNChildren(n int, typ *gti.Type, nameStub ...string) (mods, updt bool)

	// ConfigChildren configures children according to given list of
	// type-and-name's -- attempts to have minimal impact relative to existing
	// items that fit the type and name constraints (they are moved into the
	// corresponding positions), and any extra children are removed, and new
	// ones added, to match the specified config.
	// It is important that names are unique!
	//
	// IMPORTANT: returns whether any modifications were made (mods) AND if
	// that is true, the result from the corresponding UpdateStart call --
	// UpdateEnd is NOT called, allowing for further subsequent updates before
	// you call UpdateEnd(updt).
	ConfigChildren(config Config) (mods, updt bool)

	//////////////////////////////////////////////////////////////////////////
	//  Deleting Children

	// DeleteChildAtIndex deletes child at given index. It returns false
	// if there is no child at the given index. Wraps delete in UpdateStart / End
	// and sets ChildDeleted flag.
	DeleteChildAtIndex(idx int, destroy bool) bool

	// DeleteChild deletes the given child node, returning false if
	// it can not find it. Wraps delete in UpdateStart / End and
	// sets ChildDeleted flag.
	DeleteChild(child Ki, destroy bool) bool

	// DeleteChildByName deletes child node by name, returning false
	// if it can not find it. Wraps delete in UpdateStart / End and
	// sets ChildDeleted flag.
	DeleteChildByName(name string, destroy bool) bool

	// DeleteChildren deletes all children nodes -- destroy will add removed
	// children to deleted list, to be destroyed later -- otherwise children
	// remain intact but parent is nil -- could be inserted elsewhere, but you
	// better have kept a slice of them before calling this.
	DeleteChildren(destroy bool)

	// Delete deletes this node from its parent children list -- destroy will
	// add removed child to deleted list, to be destroyed later -- otherwise
	// child remains intact but parent is nil -- could be inserted elsewhere.
	Delete(destroy bool)

	// Destroy deletes and destroys all children and their childrens-children, etc.
	Destroy()

	//////////////////////////////////////////////////////////////////////////
	//  Flags

	// Is checks if flag is set, using atomic, safe for concurrent access
	Is(f enums.BitFlag) bool

	// SetFlag sets the given flag(s) to given state
	// using atomic, safe for concurrent access
	SetFlag(on bool, f ...enums.BitFlag)

	// SetChildAdded sets the ChildAdded flag -- set when notification is needed
	// for Add, Insert methods
	SetChildAdded()

	// ClearUpdateFlags resets all structure update related flags:
	// ChildAdded, ChildDeleted, ChildrenDeleted, Deleted
	// automatically called on StartUpdate to reset any old state.
	ClearUpdateFlags()

	// FlagType returns the flags of the node as the true flag type of the node,
	// which may be a type that extends the standard [Flags]. Each node type
	// that extends the flag type should define this method; for example:
	//	func (wb *WidgetBase) FlagType() enums.BitFlagSetter {
	//		return (*WidgetFlags)(&wb.Flags)
	//	}
	FlagType() enums.BitFlagSetter

	//////////////////////////////////////////////////////////////////////////
	//  Property interface with inheritance -- nodes can inherit props from parents

	// Properties (Node.Props) tell the GoGi GUI or other frameworks operating
	// on Trees about special features of each node -- functions below support
	// inheritance up Tree -- see kit convert.go for robust convenience
	// methods for converting interface{} values to standard types.
	Properties() *Props

	// SetProp sets given property key to value val -- initializes property
	// map if nil.
	SetProp(key string, val any)

	// Prop returns the property value for the given key.
	// It returns nil if it doesn't exist.
	Prop(key string) any

	// PropInherit gets property value from key with options for inheriting
	// property from parents.  If inherit, then checks all parents.
	// Returns false if not set anywhere.
	PropInherit(key string, inherit bool) (any, bool)

	// DeleteProp deletes property key on this node.
	DeleteProp(key string)

	// PropTag returns the name to look for in type properties, for types
	// that are valid options for values that can be set in Props.  For example
	// in GoGi, it is "style-props" which is then set for all types that can
	// be used in a style (colors, enum options, etc)
	PropTag() string

	//////////////////////////////////////////////////////////////////////////
	//  Tree walking and Paths
	//   note: always put function args last -- looks better for inline functions

	// WalkUp calls function on given node and all the way up to its parents,
	// and so on -- sequentially all in current go routine (generally
	// necessary for going up, which is typically quite fast anyway) -- level
	// is incremented after each step (starts at 0, goes up), and passed to
	// function -- returns false if fun aborts with false, else true.
	WalkUp(fun func(k Ki) bool) bool

	// WalkUpParent calls function on parent of node and all the way up to its
	// parents, and so on -- sequentially all in current go routine (generally
	// necessary for going up, which is typically quite fast anyway) -- level
	// is incremented after each step (starts at 0, goes up), and passed to
	// function -- returns false if fun aborts with false, else true.
	WalkUpParent(fun func(k Ki) bool) bool

	// WalkPre calls function on this node (MeFirst) and then iterates
	// in a depth-first manner over all the children.
	// The [WalkPreNode] method is called for every node, after the given function,
	// which e.g., enables nodes to also traverse additional Ki Trees (e.g., Fields),
	// including for the basic UpdateStart / End and other such infrastructure calls
	// which use WalkPre (otherwise it could just be done in the given fun).
	// The node traversal is non-recursive and uses locally-allocated state -- safe
	// for concurrent calling (modulo conflict management in function call itself).
	// Function calls are sequential all in current go routine.
	// If fun returns false then any further traversal of that branch of the tree is
	// aborted, but other branches continue -- i.e., if fun on current node
	// returns false, children are not processed further.
	WalkPre(fun func(k Ki) bool)

	// WalkPreNode is called for every node during WalkPre with the function
	// passed to WalkPre.  This e.g., enables nodes to also traverse additional
	// Ki Trees (e.g., Fields), including for the basic UpdateStart / End and
	// other such infrastructure calls.
	WalkPreNode(fun func(k Ki) bool)

	// WalkPreLevel calls function on this node (MeFirst) and then iterates
	// in a depth-first manner over all the children.
	// This version has a level var that tracks overall depth in the tree.
	// If fun returns false then any further traversal of that branch of the tree is
	// aborted, but other branches continue -- i.e., if fun on current node
	// returns false, children are not processed further.
	// Because WalkPreLevel is not used within Ki itself, it does not have its
	// own version of WalkPreNode -- that can be handled within the closure.
	WalkPreLevel(fun func(k Ki, level int) bool)

	// WalkPost iterates in a depth-first manner over the children, calling
	// doChildTestFunc on each node to test if processing should proceed (if
	// it returns false then that branch of the tree is not further processed),
	// and then calls given fun function after all of a node's children
	// have been iterated over ("Me Last").
	// This uses node state information to manage the traversal and is very fast,
	// but can only be called by one thread at a time -- use a Mutex if there is
	// a chance of multiple threads running at the same time.
	// Function calls are sequential all in current go routine.
	// The level var tracks overall depth in the tree.
	WalkPost(doChildTestFunc func(k Ki) bool, fun func(k Ki) bool)

	// WalkBreadth calls function on all children in breadth-first order
	// using the standard queue strategy.  This depends on and updates the
	// Depth parameter of the node.  If fun returns false then any further
	// traversal of that branch of the tree is aborted, but other branches continue.
	WalkBreadth(fun func(k Ki) bool)

	// UpdateStart should be called when starting to modify the tree (state or
	// structure) -- returns whether this node was first to set the Updating
	// flag (if so, all children have their Updating flag set -- pass the
	// result to UpdateEnd -- automatically determines the highest level
	// updated, within the normal top-down updating sequence -- can be called
	// multiple times at multiple levels -- it is essential to ensure that all
	// such Start's have an End!  Usage:
	//
	//   updt := n.UpdateStart()
	//   ... code
	//   n.UpdateEnd(updt)
	// or
	//   updt := n.UpdateStart()
	//   defer n.UpdateEnd(updt)
	//   ... code
	UpdateStart() bool

	// UpdateEnd should be called when done updating after an UpdateStart,
	// and passed the result of the UpdateStart call.
	// If this arg is true, the OnUpdated method will be called and the Updating
	// flag will be cleared.  Also, if any ChildDeleted flags have been set,
	// the delete manager DestroyDeleted is called.
	// If the updt bool arg is false, this function is a no-op.
	UpdateEnd(updt bool)

	//////////////////////////////////////////////////////////////////////////
	//  Deep Copy of Trees

	// CopyFrom another Ki node.  It is essential that source has Unique names!
	// The Ki copy function recreates the entire tree in the copy, duplicating
	// children etc, copying Props too.  It is very efficient by
	// using the ConfigChildren method which attempts to preserve any existing
	// nodes in the destination if they have the same name and type -- so a
	// copy from a source to a target that only differ minimally will be
	// minimally destructive.  Only copies to same types are supported.
	// Signal connections are NOT copied.  No other Ki pointers are copied,
	// and the field tag copy:"-" can be added for any other fields that
	// should not be copied (unexported, lower-case fields are not copyable).
	CopyFrom(frm Ki) error

	// Clone creates and returns a deep copy of the tree from this node down.
	// Any pointers within the cloned tree will correctly point within the new
	// cloned tree (see Copy info).
	Clone() Ki

	// CopyFieldsFrom is the base-level copy method that any copy-intensive
	// nodes should implement directly to explicitly copy relevant fields
	// that should be copied, avoiding any internal pointers etc.
	// Must explicitly call the CopyFieldsFrom method on any embedded
	// Ki types that you inherit from, and, critically, NONE of those
	// can rely on the generic Node-level version.
	CopyFieldsFrom(frm any)

	//////////////////////////////////////////////////////////////////////////
	// 	Event-specific methods

	// OnInit is called when the node is
	// initialized (ie: through InitName).
	// It is called before the node is added to the tree,
	// so it will not have any parents or siblings.
	// It will be called only once in the lifetime of the node.
	// It does nothing by default, but it can be implemented
	// by higher-level types that want to do something.
	OnInit()

	// OnAdd is called when the node is added to a parent.
	// It will be called only once in the lifetime of the node,
	// unless the node is moved. It will not be called on root
	// nodes, as they are never added to a parent.
	// It does nothing by default, but it can be implemented
	// by higher-level types that want to do something.
	OnAdd()

	// OnChildAdded is called when a node is added to
	// this node or any of its children. When a node is added to
	// a tree, it calls [OnAdd] and then this function on each of its parents,
	// going in order from the closest parent to the furthest parent.
	// This function does nothing by default, but it can be
	// implemented by higher-level types that want to do something.
	OnChildAdded(child Ki)

	// OnDelete is called when the node is deleted from a parent.
	// It will be called only once in the lifetime of the node,
	// unless the node is moved. It will not be called on root
	// nodes, as they are never deleted from a parent.
	// It does nothing by default, but it can be implemented
	// by higher-level types that want to do something.
	OnDelete()

	// OnChildDeleting is called when a node is just about to be deleted from
	// this node or any of its children. When a node is deleted from
	// a tree, it calls this function on each of its parents,
	// going in order from the closest parent to the furthest parent,
	// and then [OnDelete].
	// This function does nothing by default, but it can be
	// implemented by higher-level types that want to do something.
	OnChildDeleting(child Ki)

	// OnChildrenDeleting is called when all children are deleted from
	// this node or any of its children.
	// This function does nothing by default, but it can be
	// implemented by higher-level types that want to do something.
	OnChildrenDeleting()

	// OnUpdated is called during UpdateEnd if the updt flag is true,
	// indicating that this was the upper-most Ki node that was
	// updated in the latest round of updating.
	// This function does nothing by default, but it can be
	// implemented by higher-level types that want to do something.
	OnUpdated()
}

// see node.go for struct implementing this interface

// KiType is a Ki reflect.Type, suitable for checking for Type.Implements.
var KiType = reflect.TypeOf((*Ki)(nil)).Elem()

// todo: remove these if possible to eliminate reflect dependencies

// IsKi returns true if the given type implements the Ki interface at any
// level of embedded structure.
func IsKi(typ reflect.Type) bool {
	if typ == nil {
		return false
	}
	if typ.Implements(KiType) {
		return true
	}
	if reflect.PtrTo(typ).Implements(KiType) { // typically need the pointer type to impl
		return true
	}
	return false
}

// AsKi returns the Ki interface and Node base type for
// any given object -- if not a Ki, return values are nil.
func AsKi(v any) (Ki, *Node) {
	k, ok := v.(Ki)
	if !ok {
		return nil, nil
	}
	return k, k.AsKi()
}

// NewOfType makes a new Ki struct of given type -- must be a Ki type -- will
// return nil if not.
func NewOfType(typ *gti.Type) Ki {
	return typ.Instance.(Ki).New()
}
