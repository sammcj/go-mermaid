# Mermaid Diagram Examples

This document contains examples of all diagram types supported by mermaid-check.

## Table of Contents

1. [Flowchart](#flowchart)
2. [Sequence Diagram](#sequence-diagram)
3. [Class Diagram](#class-diagram)
4. [State Diagram](#state-diagram)
5. [Entity Relationship Diagram](#entity-relationship-diagram)
6. [Gantt Chart](#gantt-chart)
7. [Pie Chart](#pie-chart)
8. [User Journey](#user-journey)
9. [Git Graph](#git-graph)
10. [Mindmap](#mindmap)
11. [Timeline](#timeline)
12. [Quadrant Chart](#quadrant-chart)
13. [Sankey Diagram](#sankey-diagram)
14. [XY Chart](#xy-chart)
15. [C4 Diagrams](#c4-diagrams)
    - [C4 Context](#c4-context)
    - [C4 Container](#c4-container)
    - [C4 Component](#c4-component)
    - [C4 Dynamic](#c4-dynamic)
    - [C4 Deployment](#c4-deployment)

---

## Flowchart

Flowcharts visualise processes, decisions, and workflows.

```mermaid
flowchart TB
    subgraph "Agent Layer"
        Agent[Strands Agent]
    end

    subgraph "AgentCore Runtime"
        Runtime[AgentCore Runtime<br/>0.0.0.0:8000/mcp]
        Session[Session Management<br/>MCP-Session-Id header]
    end

    subgraph "MCP Server Container"
        MCP[FastMCP Server<br/>Streamable HTTP Mode]
        Tools[MCP Tools<br/>calculator, search, etc.]
    end

    Agent -->|ConverseStream API| Runtime
    Runtime -->|HTTP Streaming<br/>Accept: application/json,<br/>text/event-stream| MCP
    MCP --> Tools
    Tools -->|Real-time Events| MCP
    MCP -->|Event Stream| Runtime
    Runtime -->|Streamed Response| Agent

    Session -.->|Maintains Isolation| Runtime

    classDef llm fill:#E8F5E8,stroke:#27AE60,color:#27AE60
    classDef components fill:#F0E6FF,stroke:#8E44AD,color:#8E44AD
    classDef api fill:#FFF0E6,stroke:#E67E22,color:#E67E22

    class Agent llm
    class Runtime,Session,MCP components
    class Tools api
```

**Supported directions**: `TB` (top-bottom), `TD` (top-down), `BT` (bottom-top), `RL` (right-left), `LR` (left-right)

**Alternative syntax**: `graph` can be used instead of `flowchart`

---

## Sequence Diagram

Sequence diagrams show interactions between participants over time.

```mermaid
sequenceDiagram
    participant Alice
    participant Bob
    Alice->>John: Hello John, how are you?
    loop Healthcheck
        John->>John: Fight against hypochondria
    end
    Note right of John: Rational thoughts <br/>prevail!
    John-->>Alice: Great!
    John->>Bob: How about you?
    Bob-->>John: Jolly good!
```

**Key features**:
- Participants and actors
- Messages (solid arrows `->`, dotted arrows `-->`)
- Loops, alternatives, and parallel sections
- Notes and activation boxes

---

## Class Diagram

Class diagrams model object-oriented systems with classes, attributes, methods, and relationships.

```mermaid
classDiagram
    Animal <|-- Duck
    Animal <|-- Fish
    Animal <|-- Zebra
    Animal : +int age
    Animal : +String gender
    Animal: +isMammal()
    Animal: +mate()
    class Duck{
        +String beakColor
        +swim()
        +quack()
    }
    class Fish{
        -int sizeInFeet
        -canEat()
    }
    class Zebra{
        +bool is_wild
        +run()
    }
```

**Relationships**: `<|--` (inheritance), `*--` (composition), `o--` (aggregation), `-->` (association), `..>` (dependency)

---

## State Diagram

State diagrams model system states and transitions.

```mermaid
stateDiagram-v2
    [*] --> Still
    Still --> [*]
    Still --> Moving
    Moving --> Still
    Moving --> Crash
    Crash --> [*]
```

**Key features**:
- Initial state `[*]`
- State transitions with triggers
- Composite states
- Concurrent states with `--`

---

## Entity Relationship Diagram

ER diagrams model database schemas and relationships between entities.

```mermaid
erDiagram
    CUSTOMER ||--o{ ORDER : places
    ORDER ||--|{ LINE-ITEM : contains
    CUSTOMER }|..|{ DELIVERY-ADDRESS : uses
```

**Cardinality notation**:
- `||` exactly one
- `o{` zero or more
- `|{` one or more
- `}|` one or more (right side)

---

## Gantt Chart

Gantt charts display project schedules and task dependencies.

```mermaid
gantt
    title A Gantt Diagram
    dateFormat YYYY-MM-DD
    section Section
        A task          :a1, 2014-01-01, 30d
        Another task    :after a1, 20d
    section Another
        Task in Another :2014-01-12, 12d
        another task    :24d
```

**Key features**:
- Date formats and ranges
- Task dependencies
- Sections for grouping
- Active, done, and critical task states

---

## Pie Chart

Pie charts show proportional data distribution.

```mermaid
pie title Pets adopted by volunteers
    "Dogs" : 386
    "Cats" : 85
    "Rats" : 15
```

**Features**:
- Optional title
- Label and value pairs
- Automatic percentage calculation

---

## User Journey

User journey diagrams map user experiences and satisfaction levels.

```mermaid
journey
    title My working day
    section Go to work
      Make tea: 5: Me
      Go upstairs: 3: Me
      Do work: 1: Me, Cat
    section Go home
      Go downstairs: 5: Me
      Sit down: 5: Me
```

**Features**:
- Satisfaction scores (1-5)
- Multiple actors
- Sectioned activities

---

## Git Graph

Git graph diagrams visualise Git branching and merging workflows.

```mermaid
gitGraph
    commit
    commit
    branch develop
    checkout develop
    commit
    commit
    checkout main
    merge develop
    commit
```

**Key commands**:
- `commit` - create commit
- `branch` - create branch
- `checkout` - switch branch
- `merge` - merge branches
- `cherry-pick` - cherry-pick commit

---

## Mindmap

Mindmaps organise ideas hierarchically.

```mermaid
mindmap
  root<br/>My Mindmap
    Origins
      Long history
      Popularisation
        British popular psychology author Tony Buzan
    Research
      On effectiveness<br/>and features
      On Automatic creation
        Uses
          Creative techniques
          Strategic planning
          Argument mapping
    Tools
      Pen and paper
      Mermaid
```

**Features**:
- Hierarchical structure through indentation
- Root node with branches
- Multiple levels of nesting
- Line breaks with `<br/>`

---

## Timeline

Timeline diagrams display chronological events.

```mermaid
timeline
    title History of Social Media Platform
    2002 : LinkedIn
    2004 : Facebook
         : Google
    2005 : Youtube
    2006 : Twitter
```

**Features**:
- Chronological ordering
- Multiple events per time period
- Optional title

---

## Quadrant Chart

Quadrant charts plot items in a 2D space with labelled quadrants.

```mermaid
quadrantChart
    title Reach and engagement of campaigns
    x-axis Low Reach --> High Reach
    y-axis Low Engagement --> High Engagement
    quadrant-1 We should expand
    quadrant-2 Need to promote
    quadrant-3 Re-evaluate
    quadrant-4 May be improved
    Campaign A: [0.3, 0.6]
    Campaign B: [0.45, 0.23]
    Campaign C: [0.57, 0.69]
    Campaign D: [0.78, 0.34]
    Campaign E: [0.40, 0.34]
    Campaign F: [0.35, 0.78]
```

**Features**:
- Custom axis labels
- Four quadrant labels
- Coordinate-based data points (0-1 range)

---

## Sankey Diagram

Sankey diagrams visualise flows and quantities between nodes.

```mermaid
sankey-beta

Agricultural 'waste',Bio-conversion,124.729
Bio-conversion,Liquid,0.597
Bio-conversion,Losses,26.862
Bio-conversion,Solid,280.322
Bio-conversion,Gas,81.144
```

**Features**:
- Source and target nodes
- Flow quantities
- Multi-level flow chains
- Currently in beta (`sankey-beta`)

---

## XY Chart

XY charts display data series on X-Y axes.

```mermaid
xychart-beta
    title "Sales Revenue"
    x-axis [jan, feb, mar, apr, may, jun, jul, aug, sep, oct, nov, dec]
    y-axis "Revenue (in $)" 4000 --> 11000
    bar [5000, 6000, 7500, 8200, 9500, 10500, 11000, 10200, 9200, 8500, 7000, 6000]
    line [5000, 6000, 7500, 8200, 9500, 10500, 11000, 10200, 9200, 8500, 7000, 6000]
```

**Features**:
- Bar and line series
- Custom axis ranges
- Multiple data series
- Currently in beta (`xychart-beta`)

---

## C4 Diagrams

C4 diagrams model software architecture at different levels of abstraction.

### C4 Context

System context diagrams show how your system fits in the world.

```mermaid
C4Context
    title System Context diagram for Internet Banking System

    Person(customerA, "Banking Customer A", "A customer of the bank")
    Person(customerB, "Banking Customer B")
    Person_Ext(customerC, "Banking Customer C", "External customer")
    System(SystemAA, "Internet Banking System", "Allows customers to view information about their bank accounts")

    System_Ext(SystemE, "Mainframe Banking System", "Stores all customer information")

    Rel(customerA, SystemAA, "Uses")
    Rel(SystemAA, SystemE, "Uses")
    Rel(customerB, SystemAA, "Uses")
    Rel(customerC, SystemAA, "Uses")
```

### C4 Container

Container diagrams show the high-level technology choices.

```mermaid
C4Container
    title Container diagram for Internet Banking System

    Person(customer, "Banking Customer", "A customer of the bank")
    System_Boundary(c1, "Internet Banking System") {
        Container(web_app, "Web Application", "Java, Spring MVC", "Delivers content")
        Container(spa, "Single-Page App", "JavaScript, Angular", "Provides banking functionality")
        Container(mobile_app, "Mobile App", "C#, Xamarin", "Provides limited banking functionality")
        ContainerDb(database, "Database", "SQL Database", "Stores user data")
    }
    System_Ext(email_system, "E-Mail System", "External e-mail system")

    Rel(customer, web_app, "Uses", "HTTPS")
    Rel(customer, spa, "Uses", "HTTPS")
    Rel(customer, mobile_app, "Uses")
    Rel(web_app, spa, "Delivers")
    Rel(spa, database, "Reads/Writes", "SQL/TCP")
    Rel(mobile_app, database, "Reads/Writes", "SQL/TCP")
    Rel(database, email_system, "Sends email", "SMTP")
```

### C4 Component

Component diagrams show how a container is made up of components.

```mermaid
C4Component
    title Component diagram for Internet Banking System - API Application

    Container(spa, "Single Page Application", "JavaScript, Angular", "Provides banking functionality")
    Container(ma, "Mobile App", "C#, Xamarin", "Provides limited banking functionality")
    ContainerDb(db, "Database", "SQL Database", "Stores user data")

    Container_Boundary(api, "API Application") {
        Component(sign, "Sign In Controller", "MVC Rest Controller", "Allows users to sign in")
        Component(accounts, "Accounts Summary Controller", "MVC Rest Controller", "Provides customers account summary")
        Component(security, "Security Component", "Spring Bean", "Provides authentication and authorisation")
        Component(mbsfacade, "Mainframe Banking System Facade", "Spring Bean", "A facade onto the mainframe system")
    }

    Rel(sign, security, "Uses")
    Rel(accounts, mbsfacade, "Uses")
    Rel(security, db, "Reads/Writes", "SQL/TCP")
    Rel(spa, sign, "Uses", "JSON/HTTPS")
    Rel(spa, accounts, "Uses", "JSON/HTTPS")
    Rel(ma, sign, "Uses", "JSON/HTTPS")
    Rel(ma, accounts, "Uses", "JSON/HTTPS")
```

### C4 Dynamic

Dynamic diagrams show how components collaborate for a specific scenario.

```mermaid
C4Dynamic
    title Dynamic diagram for Internet Banking System - API Application

    ContainerDb(c4, "Database", "SQL Database", "Stores user data")
    Container(c1, "Single Page Application", "JavaScript, Angular", "Provides banking functionality")
    Container_Boundary(b, "API Application") {
      Component(c3, "Security Component", "Spring Bean", "Provides authentication")
      Component(c2, "Sign In Controller", "MVC Rest Controller", "Allows users to sign in")
    }

    Rel(c1, c2, "Submits credentials", "JSON/HTTPS")
    Rel(c2, c3, "Validates credentials")
    Rel(c3, c4, "Reads user info", "SQL/TCP")

    UpdateRelStyle(c1, c2, $offsetY="60", $offsetX="90")
    UpdateRelStyle(c2, c3, $offsetX="-40", $offsetY="60")
    UpdateRelStyle(c3, c4, $offsetY="-40", $offsetX="-40")
```

### C4 Deployment

Deployment diagrams show how containers are deployed to infrastructure.

```mermaid
C4Deployment
    title Deployment diagram for Internet Banking System

    Deployment_Node(plc, "Customer's Computer", "Windows, macOS, Linux"){
        Deployment_Node(browser, "Web Browser", "Chrome, Firefox, Safari"){
            Container(spa, "Single-Page App", "JavaScript, Angular", "Provides banking functionality")
        }
    }

    Deployment_Node(aws, "Amazon Web Services"){
        Deployment_Node(ec2, "EC2 Instance"){
            Container(web, "Web Application", "Java, Spring MVC", "Delivers content")
        }
        Deployment_Node(rds, "RDS Instance"){
            ContainerDb(db, "Database", "SQL Database", "Stores user data")
        }
    }

    Rel(spa, web, "Uses", "HTTPS")
    Rel(web, db, "Reads/Writes", "SQL/TCP")
```

---

## Notes

- All examples use British English spelling in comments and labels where applicable
- Line breaks in labels use `<br/>` not `\n`
- Comments in Mermaid use `%%` prefix
- Round brackets `()` should be avoided in labels; use brackets `[]` or braces `{}` for node shapes instead
