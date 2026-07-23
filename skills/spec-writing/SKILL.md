---
name: spec-writing
description: >
   Write, revise, reconcile, or review detailed software specifications.
   Explains prose style, technical language, level of detail, and notation.
---

# Natural Language Software Specification

When asked for a spec you shall use the "nlspec" (Natural Language Spec) style,
described here.

## Fundamentals

A nlspec is a minimally spanning document providing enough information for a
frontier coding agent to build the described software.  At its heart it contains
the following conceptual.

* Overview and Goals -- a high-level description of the software, the needs it
  solves, and key design and esthetic principles.
* Definition of done -- The list of behavioral claims that must be demonstrated
  working, to accept the software as ready.
* Exclusions (not-building) -- Considerations that are explicitly excluded from
  the design, and agents are advised/enjoined not to build them, to prevent
  over-engineering and undue complexity.
* Central data types and algorithms, further described below. Use
  language-neutral pseudocode except when specifying public syntax or wire
  formats that must be shown exactly.
* Seams, interfaces, contracts.  Use a combination of prose and pseudocode to
  divide the project into components. Each component must have an explicit
  responsibility, boundary, and contract.

## Scope

A nlspec is NOT a technical specification in the usual sense.  Rather than
specifying EVERYTHING, it specifies only the essential structures, behaviors,
and considerations needed for an agent to be able to build the described software.

**Minimalism Requirement** -- Omit interchangeable implementation details.
Include externally visible behavior, invariants, defaults, ordering rules, and
boundary decisions that reasonable implementations could interpret differently.

The seams should be described in enough detail that each component can be
independently implemented and integrated through its stated contracts.

## Workflow

1. Read the complete existing spec and its referenced specs. Identify the
   authoritative source and whether the task is drafting, revising,
   reconciling, or reviewing.
2. Define the goals, exclusions, architecture, central types, contracts, and
   observable definition of done. Where relevant, specify defaults,
   precedence, lifecycle, failure, concurrency, persistence, and observability.
3. Treat prose as the canonical contract. Use tables, grammars, state machines,
   and pseudocode only when they remove ambiguity.
4. Re-read the complete result. Check terminology and every repeated rule
   across prose, tables, examples, appendices, related specs, and the definition
   of done. Each acceptance claim must trace to specified behavior.
5. Report changed decisions, unresolved questions, verification performed, and
   remaining proof gaps.

## Technical Artifacts

Do not prescribe an implementation language. Use pseudocode like the following
instead. Exact examples of a specified language, schema, or wire format are
allowed. Seek HITL approval before adding a materially new public concept or
contract.

### Record

Describes a data type definition.

```
RECORD Session:
    id                : String                  -- UUID, assigned at creation
    provider_profile  : ProviderProfile         -- tools + system prompt for the active model
    execution_env     : ExecutionEnvironment    -- where tools run
    history           : List<Turn>              -- ordered conversation turns
    event_emitter     : EventEmitter            -- delivers events to host application
    config            : SessionConfig           -- limits, timeouts, settings
    state             : SessionState            -- current lifecycle state
    llm_client        : Client                  -- from the Unified LLM SDK
    steering_queue    : Queue<String>           -- messages to inject between tool rounds
    followup_queue    : Queue<String>           -- messages to process after current input completes
    subagents         : Map<String, SubAgent>   -- active child agents
```

### Enum

```
ENUM SessionState:
    IDLE              -- waiting for user input
    PROCESSING        -- running the agentic loop
    AWAITING_INPUT    -- model asked the user a question
    CLOSED            -- session terminated (normal or error)
```

### Interface

Preferred way to define seams.  Should include type definitions for user-defined
types.

```
RECORD ToolDefinition:
    name        : String            -- unique identifier
    description : String            -- for the LLM
    parameters  : Dict              -- JSON Schema (root must be "object")

RECORD RegisteredTool:
    definition  : ToolDefinition
    executor    : Function          -- (arguments, execution_env) -> String

RECORD ToolRegistry:
    _tools      : Map<String, RegisteredTool>

    register(tool)                  -- add or replace a tool
    unregister(name)                -- remove a tool
    get(name) -> RegisteredTool | None
    definitions() -> List<ToolDefinition>
    names() -> List<String>

INTERFACE ProviderProfile:
    id              : String            -- "openai", "anthropic", "gemini"
    model           : String            -- model identifier (e.g., "gpt-5.2-codex")
    tool_registry   : ToolRegistry      -- all tools available to this profile

    FUNCTION build_system_prompt(environment, project_docs) -> String
    FUNCTION tools() -> List<ToolDefinition>
    FUNCTION provider_options() -> Map | None

    -- Capability flags
    supports_reasoning           : Boolean
    supports_streaming           : Boolean
    supports_parallel_tool_calls : Boolean
    context_window_size          : Integer
```

### State Machine

Coupled with a state enum, to define the allowed transitions.

```
IDLE -> PROCESSING          -- on submit()
PROCESSING -> PROCESSING    -- tool loop continues
PROCESSING -> AWAITING_INPUT -- model asks user a question (no tool calls, open-ended)
PROCESSING -> IDLE          -- natural completion or turn limit
PROCESSING -> CLOSED        -- unrecoverable error
IDLE -> CLOSED              -- explicit close()
any -> CLOSED               -- abort signal (after graceful shutdown cleanup)
AWAITING_INPUT -> PROCESSING -- user provides answer
```

### Function

A pseudocode algorithm definition. Use only for core behavior whose ordering,
precedence, or edge cases would be ambiguous in prose. It defines required
semantics, not a prescribed implementation.

```
FUNCTION check_context_usage(session):
    approx_tokens = total_chars_in_history(session.history) / 4
    threshold = session.provider_profile.context_window_size * 0.8
    IF approx_tokens > threshold:
        session.emit(WARNING, message = "Context usage at ~"
            + ROUND(approx_tokens / session.provider_profile.context_window_size * 100)
            + "% of context window")
```

### Service Definition

This should be rare as well, only to be used when the algorithms being shown
are non-obvious but simple and/or the prose description of the behavior
would be more verbose.

```
HandlerRegistry:
    handlers        : Map<String, Handler>   -- type string -> handler instance
    default_handler : Handler                -- fallback handler (typically codergen)

    FUNCTION register(type_string, handler):
        handlers[type_string] = handler
        -- Registering for an already-registered type replaces the previous handler

    FUNCTION resolve(node) -> Handler:
        -- 1. Explicit type attribute
        IF node.type is not empty AND node.type IN handlers:
            RETURN handlers[node.type]

        -- 2. Shape-based resolution
        handler_type = SHAPE_TO_TYPE[node.shape]
        IF handler_type IN handlers:
            RETURN handlers[handler_type]

        -- 3. Default
        RETURN default_handler
```
