package vm

// nop
//
//    nop
//    ()
//    ()
//    ()
func (vm *Vm) nop() {}

// Get local variable (pointed by idx' and level').
// 'level' indicates the nesting depth from the current block.
//
//    getlocal
//    (lindex_t idx, rb_num_t level)
//    ()
//    (VALUE val)
func (vm *Vm) getlocal() {}

// Set a local variable (pointed to by 'idx') as val.
// 'level' indicates the nesting depth from the current block.
//
//    setlocal
//    (lindex_t idx, rb_num_t level)
//    (VALUE val)
//    ()
func (vm *Vm) setlocal() {}

// Get value of special local variable ($~, $_, ..).
//
//    getspecial
//    (rb_num_t key, rb_num_t type)
//    ()
//    (VALUE val)
func (vm *Vm) getspecial() {}

// Set value of special local variable ($~, $_, ...) to obj.
//
//    setspecial
//    (rb_num_t key)
//    (VALUE obj)
//    ()
func (vm *Vm) setspecial() {}

// Get value of instance variable id of self.
// If is_local is not 0, get value of class local variable.
//
//    getinstancevariable
//    (ID id, IC ic)
//    ()
//    (VALUE val)
func (vm *Vm) getinstancevariable() {}

// Set value of instance variable id of self to val.
// If is_local is not 0, set value of class local variable.
//
//    setinstancevariable
//    (ID id, IC ic)
//    (VALUE val)
//    ()
func (vm *Vm) setinstancevariable() {}

// Get value of class variable id of klass as val.
//
//    getclassvariable
//    (ID id)
//    ()
//    (VALUE val)
func (vm *Vm) getclassvariable() {}

// Set value of class variable id of klass as val.
//
//    setclassvariable
//    (ID id)
//    (VALUE val)
//    ()
func (vm *Vm) setclassvariable() {}

// Get constant variable id. If klass is Qnil, constants
// are searched in the current scope. If klass is Qfalse, constants
// are searched as top level constants. Otherwise, get constant under klass
// class or module.
//
//    getconstant
//    (ID id)
//    (VALUE klass)
//    (VALUE val)
func (vm *Vm) getconstant() {}

// Set constant variable id. If klass is Qfalse, constant
// is able to access in this scope. if klass is Qnil, set
// top level constant. otherwise, set constant under klass
// class or module.
//
//    setconstant
//    (ID id)
//    (VALUE val, VALUE cbase)
//    ()
func (vm *Vm) setconstant() {}

// get global variable id.
//
//    getglobal
//    (GENTRY entry)
//    ()
//    (VALUE val)
func (vm *Vm) getglobal() {}

// set global variable id as val.
//
//    setglobal
//    (GENTRY entry)
//    (VALUE val)
//    ()
func (vm *Vm) setglobal() {}

// put nil to stack.
//
//    putnil
//    ()
//    ()
//    (VALUE val)
func (vm *Vm) putnil() {}

// put self.
//
//    putself
//    ()
//    ()
//    (VALUE val)
func (vm *Vm) putself() {}

// put some object.
// i.e. Fixnum, true, false, nil, and so on.
//
//    putobject
//    (VALUE val)
//    ()
//    (VALUE val)
func (vm *Vm) putobject() {}

// put special object."value_type" is for expansion.
//
//    putspecialobject
//    (rb_num_t value_type)
//    ()
//    (VALUE val)
func (vm *Vm) putspecialobject() {}

// put iseq value.
//
//    putiseq
//    (ISEQ iseq)
//    ()
//    (VALUE ret)
func (vm *Vm) putiseq() {}

// put string val. string will be copied.
//
//    putstring
//    (VALUE str)
//    ()
//    (VALUE val)
func (vm *Vm) putstring() {}

// put concatenate strings
//
//    concatstrings
//    (rb_num_t num)
//    (...)
//    (VALUE val) // inc += 1 - num;
func (vm *Vm) concatstrings() {}

// to_str
//
//    tostring
//    ()
//    (VALUE val)
//    (VALUE val)
func (vm *Vm) tostring() {}

// to Regexp
//
//    toregexp
//    (rb_num_t opt, rb_num_t cnt)
//    (...)
//    (VALUE val) // inc += 1 - cnt;
func (vm *Vm) toregexp() {}

// put new array.
//
//    newarray
//    (rb_num_t num)
//    (...)
//    (VALUE val) // inc += 1 - num;
func (vm *Vm) newarray() {}

// dup array
//
//    duparray
//    (VALUE ary)
//    ()
//    (VALUE val)
func (vm *Vm) duparray() {}

// expand array to num objects.
//
//    expandarray
//    (rb_num_t num, rb_num_t flag)
//    (..., VALUE ary)
//    (...) // inc += num - 1 + (flag & 1 ? 1 : 0);
func (vm *Vm) expandarray() {}

// concat two arrays
//
//    concatarray
//    ()
//    (VALUE ary1, VALUE ary2st)
//    (VALUE ary)
func (vm *Vm) concatarray() {}

// splat array
//
//    splatarray
//    (VALUE flag)
//    (VALUE ary)
//    (VALUE obj)
func (vm *Vm) splatarray() {}

// put new Hash.
//
//    newhash
//    (rb_num_t num)
//    (...)
//    (VALUE val) // inc += 1 - num;
func (vm *Vm) newhash() {}

// put new Range object.(Range.new(low, high, flag))
//
//    newrange
//    (rb_num_t flag)
//    (VALUE low, VALUE high)
//    (VALUE val)
func (vm *Vm) newrange() {}

// pop from stack.
//
//    pop
//    ()
//    (VALUE val)
//    ()
func (vm *Vm) pop() {}

// duplicate stack top.
//
//    dup
//    ()
//    (VALUE val)
//    (VALUE val1, VALUE val2)
func (vm *Vm) dup() {}

// duplicate stack top n elements
//
//    dupn
//    (rb_num_t n)
//    (...)
//    (...) // inc += n;
func (vm *Vm) dupn() {}

// swap top 2 vals
//
//    swap
//    ()
//    (VALUE val, VALUE obj)
//    (VALUE obj, VALUE val)
func (vm *Vm) swap() {}

// for stack caching.
//
//    reput
//    ()
//    (..., VALUE val)
//    (VALUE val) // inc += 0;
func (vm *Vm) reput() {}

// get nth stack value from stack top
//
//    topn
//    (rb_num_t n)
//    (...)
//    (VALUE val) // inc += 1;
func (vm *Vm) topn() {}

// set Nth stack entry to stack top
//
//    setn
//    (rb_num_t n)
//    (..., VALUE val)
//    (VALUE val) // inc += 0
func (vm *Vm) setn() {}

// empt current stack
//
//    adjuststack
//    (rb_num_t n)
//    (...)
//    (...) // inc -= n
func (vm *Vm) adjuststack() {}

// defined?
//
//    defined
//    (rb_num_t op_type, VALUE obj, VALUE needstr)
//    (VALUE v)
//    (VALUE val)
func (vm *Vm) defined() {}

// check target' matches pattern'.
// flag & VM_CHECKMATCH_TYPE_MASK' describe how to check pattern.
// VM_CHECKMATCH_TYPE_WHEN: ignore target and check pattern is truthy.
// VM_CHECKMATCH_TYPE_CASE: check patten === target'.
// VM_CHECKMATCH_TYPE_RESCUE: check pattern.kind_op?(Module) && pattern == target'.
// if flag & VM_CHECKMATCH_ARRAY' is not 0, then patten' is array of patterns.
//
//    checkmatch
//    (rb_num_t flag)
//    (VALUE target, VALUE pattern)
//    (VALUE result)
func (vm *Vm) checkmatch() {}

// trace
//
//    trace
//    (rb_num_t nf)
//    ()
//    ()
func (vm *Vm) trace() {}

// enter class definition scope. if super is Qfalse, and clsas
// "klass" is defined, it's redefine. otherwise, define "klass" class.
//
//    defineclass
//    (ID id, ISEQ class_iseq, rb_num_t flags)
//    (VALUE cbase, VALUE super)
//    (VALUE val)
func (vm *Vm) defineclass() {}

// invoke method.
//
//    send
//    (CALL_INFO ci)
//    (...)
//    (VALUE val) // inc += - (int)(ci->orig_argc + ((ci->flag & VM_CALL_ARGS_BLOCKARG) ? 1 : 0));
func (vm *Vm) send() {}

// Freeze a string
//
//    opt_str_freeze
//    (VALUE str)
//    ()
//    (VALUE val)
func (vm *Vm) opt_str_freeze() {}

// Invoke method without block, splat
//
//    opt_send_simple
//    (CALL_INFO ci)
//    (...)
//    (VALUE val) // inc += -ci->orig_argc;
func (vm *Vm) opt_send_simple() {}

// super(args) # args.size => num
//
//    invokesuper
//    (CALL_INFO ci)
//    (...)
//    (VALUE val) // inc += - (int)(ci->orig_argc + ((ci->flag & VM_CALL_ARGS_BLOCKARG) ? 1 : 0));
func (vm *Vm) invokesuper() {}

// yield(args)
//
//    invokeblock
//    (CALL_INFO ci)
//    (...)
//    (VALUE val)  // inc += 1 - ci->orig_argc;
func (vm *Vm) invokeblock() {}

// return from this scope.
//
//    leave
//    ()
//    (VALUE val)
//    (VALUE val)
func (vm *Vm) leave() {}

// longjump
//
//    throw
//    (rb_num_t throw_state)
//    (VALUE throwobj)
//    (VALUE val)
func (vm *Vm) throw() {}

// set PC to (PC + dst).
//
//    jump
//    (OFFSET dst)
//    ()
//    ()
func (vm *Vm) jump() {}

// if val is not false or nil, set PC to (PC + dst).
//
//    branchif
//    (OFFSET dst)
//    (VALUE val)
//    ()
func (vm *Vm) branchif() {}

// if val is false or nil, set PC to (PC + dst).
//
//    branchunless
//    (OFFSET dst)
//    (VALUE val)
//    ()
func (vm *Vm) branchunless() {}

// push inline-cached value and go to dst if it is valid
//
//    getinlinecache
//    (OFFSET dst, IC ic)
//    ()
//    (VALUE val)
func (vm *Vm) getinlinecache() {}

// set inline cache
//
//    setinlinecache
//    (IC ic)
//    (VALUE val)
//    (VALUE val)
func (vm *Vm) setinlinecache() {}

// run iseq only once
//
//    once
//    (ISEQ iseq, IC ic)
//    ()
//    (VALUE val)
func (vm *Vm) once() {}

// case dispatcher, jump by table if possible
//
//    opt_case_dispatch
//    (CDHASH hash, OFFSET else_offset)
//    (..., VALUE key)
//    () // inc += -1;
func (vm *Vm) opt_case_dispatch() {}

// optimized X+Y.
//
//    opt_plus
//    (CALL_INFO ci)
//    (VALUE recv, VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_plus() {}

// optimized X-Y.
//
//    opt_minus
//    (CALL_INFO ci)
//    (VALUE recv, VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_minus() {}

// optimized X*Y.
//
//    opt_mult
//    (CALL_INFO ci)
//    (VALUE recv, VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_mult() {}

// optimized X/Y.
//
//    opt_div
//    (CALL_INFO ci)
//    (VALUE recv, VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_div() {}

// optimized X%Y.
//
//    opt_mod
//    (CALL_INFO ci)
//    (VALUE recv, VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_mod() {}

// optimized X==Y.
//
//    opt_eq
//    (CALL_INFO ci)
//    (VALUE recv, VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_eq() {}

// optimized X!=Y.
//
//    opt_neq
//    (CALL_INFO ci, CALL_INFO ci_eq)
//    (VALUE recv, VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_neq() {}

// optimized X<Y.
//
//    opt_lt
//    (CALL_INFO ci)
//    (VALUE recv, VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_lt() {}

// optimized X<=Y.
//
//    opt_le
//    (CALL_INFO ci)
//    (VALUE recv, VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_le() {}

// optimized X>Y.
//
//    opt_gt
//    (CALL_INFO ci)
//    (VALUE recv, VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_gt() {}

// optimized X>=Y.
//
//    opt_ge
//    (CALL_INFO ci)
//    (VALUE recv, VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_ge() {}

// <<
//
//    opt_ltlt
//    (CALL_INFO ci)
//    (VALUE recv, VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_ltlt() {}

// []
//
//    opt_aref
//    (CALL_INFO ci)
//    (VALUE recv, VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_aref() {}

// recv[obj] = set
//
//    opt_aset
//    (CALL_INFO ci)
//    (VALUE recv, VALUE obj, VALUE set)
//    (VALUE val)
func (vm *Vm) opt_aset() {}

// optimized length
//
//    opt_length
//    (CALL_INFO ci)
//    (VALUE recv)
//    (VALUE val)
func (vm *Vm) opt_length() {}

// optimized size
//
//    opt_size
//    (CALL_INFO ci)
//    (VALUE recv)
//    (VALUE val)
func (vm *Vm) opt_size() {}

// optimized empty?
//
//    opt_empty_p
//    (CALL_INFO ci)
//    (VALUE recv)
//    (VALUE val)
func (vm *Vm) opt_empty_p() {}

// optimized succ
//
//    opt_succ
//    (CALL_INFO ci)
//    (VALUE recv)
//    (VALUE val)
func (vm *Vm) opt_succ() {}

// optimized not
//
//    opt_not
//    (CALL_INFO ci)
//    (VALUE recv)
//    (VALUE val)
func (vm *Vm) opt_not() {}

// optimized regexp match
//
//    opt_regexpmatch1
//    (VALUE r)
//    (VALUE obj)
//    (VALUE val)
func (vm *Vm) opt_regexpmatch1() {}

// optimized regexp match 2
//
//    opt_regexpmatch2
//    (CALL_INFO ci)
//    (VALUE obj2, VALUE obj1)
//    (VALUE val)
func (vm *Vm) opt_regexpmatch2() {}

// call native compiled method
//
//    opt_call_c_function
//    (rb_insn_func_t funcptr)
//    ()
//    ()
func (vm *Vm) opt_call_c_function() {}

// BLT
//
//    bitblt
//    ()
//    ()
//    (VALUE ret)
func (vm *Vm) bitblt() {}

// The Answer to Life, the Universe, and Everything
//
//    answer
//    ()
//    ()
//    (VALUE ret)
func (vm *Vm) answer() {}
