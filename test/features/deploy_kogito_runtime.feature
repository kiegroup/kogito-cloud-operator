# Disabled until we can integrate it as part of the pipeline, see KOGITO-2278
@disabled
Feature: Deploy Kogito Runtime

  Background:
    Given Namespace is created

  @smoke
  Scenario Outline: Deploy Simplest Scenario using Kogito Runtime
    Given Kogito Operator is deployed
    And Clone Kogito examples into local directory
    And Local example service "<example-service>" is built by Maven using profile "<profile>"
    And Local example service "<example-service>" is deployed to image registry, image tag stored as variable "built-image"

    When Deploy <runtime> example service using image in variable "built-image" with configuration:
      | config | persistence | disabled |
    
    Then Kogito Runtime "<example-service>" has 1 pods running within 10 minutes
    And Service "<example-service>" with process name "orders" is available within 2 minutes

    @springboot
    Examples:
      | runtime    | example-service            | profile |
      | springboot | process-springboot-example | default |

    @quarkus
    Examples:
      | runtime    | example-service         | profile |
      | quarkus    | process-quarkus-example | default |

    @quarkus
    @native
    Examples:
      | runtime    | example-service         | profile |
      | quarkus    | process-quarkus-example | native  |

#####

  @persistence
  Scenario Outline: Deploy Scenario with persistence using Kogito Runtime
    Given Kogito Operator is deployed with Infinispan operator
    And Clone Kogito examples into local directory
    And Local example service "<example-service>" is built by Maven using profile "<profile>"
    And Local example service "<example-service>" is deployed to image registry, image tag stored as variable "built-image"

    When Deploy <runtime> example service using image in variable "built-image" with configuration:
      | config | persistence | enabled |
    And Kogito Runtime "<example-service>" has 1 pods running within 10 minutes
    And Start "orders" process on service "<example-service>" within 3 minutes with body:
      """json
      {
        "approver" : "john", 
        "order" : {
          "orderNumber" : "12345", 
          "shipped" : false
        }
      }
      """
    
    Then Service "<example-service>" contains 1 instances of process with name "orders"

    When Scale Kogito Runtime "<example-service>" to 0 pods within 2 minutes
    And Scale Kogito Runtime "<example-service>" to 1 pods within 2 minutes

    Then Service "<example-service>" contains 1 instances of process with name "orders" within 2 minutes

    @springboot
    Examples:
      | runtime    | example-service            | profile     |
      | springboot | process-springboot-example | persistence |

    @quarkus
    Examples:
      | runtime    | example-service         | profile     |
      | quarkus    | process-quarkus-example | persistence |

    @quarkus
    @native
    Examples:
      | runtime    | example-service         | profile            |
      | quarkus    | process-quarkus-example | native,persistence |

#####

  @events
  Scenario Outline: Deploy Scenario with events using Kogito Runtime
    Given Kogito Operator is deployed with Infinispan and Kafka operators
    And Install Kogito Data Index with 1 replicas
    And Clone Kogito examples into local directory
    And Local example service "<example-service>" is built by Maven using profile "<profile>"
    And Local example service "<example-service>" is deployed to image registry, image tag stored as variable "built-image"

    When Deploy <runtime> example service using image in variable "built-image" with configuration:
      | config | persistence | enabled |
      | config | events      | enabled  |
    And Kogito Runtime "<example-service>" has 1 pods running within 10 minutes
    And Start "orders" process on service "<example-service>" within 3 minutes with body:
      """json
      {
        "approver" : "john", 
        "order" : {
          "orderNumber" : "12345", 
          "shipped" : false
        }
      }
      """
    
    Then GraphQL request on Data Index service returns ProcessInstances processName "orders" within 2 minutes

    @springboot
    Examples:
      | runtime    | example-service            | profile            |
      | springboot | process-springboot-example | persistence,events |

    @quarkus
    Examples:
      | runtime    | example-service         | profile            |
      | quarkus    | process-quarkus-example | persistence,events |

    @quarkus
    @native
    Examples:
      | runtime    | example-service         | profile                   |
      | quarkus    | process-quarkus-example | native,persistence,events |