<!--
**Note:** When your KEP is complete, all of these comment blocks should be removed.

To get started with this template:

- [ ] **Pick a hosting SIG.**
  Make sure that the problem space is something the SIG is interested in taking
  up. KEPs should not be checked in without a sponsoring SIG.
- [ ] **Create an issue in kubernetes/enhancements**
  When filing an enhancement tracking issue, please make sure to complete all
  fields in that template. One of the fields asks for a link to the KEP. You
  can leave that blank until this KEP is filed, and then go back to the
  enhancement and add the link.
- [ ] **Make a copy of this template directory.**
  Copy this template into the owning SIG's directory and name it
  `NNNN-short-descriptive-title`, where `NNNN` is the issue number (with no
  leading-zero padding) assigned to your enhancement above.
- [ ] **Fill out as much of the kep.yaml file as you can.**
  At minimum, you should fill in the "Title", "Authors", "Owning-sig",
  "Status", and date-related fields.
- [ ] **Fill out this file as best you can.**
  At minimum, you should fill in the "Summary" and "Motivation" sections.
  These should be easy if you've preflighted the idea of the KEP with the
  appropriate SIG(s).
- [ ] **Create a PR for this KEP.**
  Assign it to people in the SIG who are sponsoring this process.
- [ ] **Merge early and iterate.**
  Avoid getting hung up on specific details and instead aim to get the goals of
  the KEP clarified and merged quickly. The best way to do this is to just
  start with the high-level sections and fill out details incrementally in
  subsequent PRs.

Just because a KEP is merged does not mean it is complete or approved. Any KEP
marked as `provisional` is a working document and subject to change. You can
denote sections that are under active debate as follows:

```
<<[UNRESOLVED optional short context or usernames ]>>
Stuff that is being argued.
<<[/UNRESOLVED]>>
```

When editing KEPS, aim for tightly-scoped, single-topic PRs to keep discussions
focused. If you disagree with what is already in a document, open a new PR
with suggested changes.

One KEP corresponds to one "feature" or "enhancement" for its whole lifecycle.
You do not need a new KEP to move from beta to GA, for example. If
new details emerge that belong in the KEP, edit the KEP. Once a feature has become
"implemented", major changes should get new KEPs.

The canonical place for the latest set of instructions (and the likely source
of this file) is [here](/keps/NNNN-kep-template/README.md).

**Note:** Any PRs to move a KEP to `implementable`, or significant changes once
it is marked `implementable`, must be approved by each of the KEP approvers.
If none of those approvers are still appropriate, then changes to that list
should be approved by the remaining approvers and/or the owning SIG (or
SIG Architecture for cross-cutting KEPs).
-->
# KEP-NNNN: Your short, descriptive title
<!--DIRECTIVE: required: true-->
<!--
This is the title of your KEP. Keep it short, simple, and descriptive. A good
title can help communicate what the KEP is and should be considered as part of
any review.
-->

<!--
A table of contents is helpful for quickly jumping to sections of a KEP and for
highlighting any additional information provided beyond the standard KEP
template.

Ensure the TOC is wrapped with
  <code>&lt;!-- toc --&rt;&lt;!-- /toc --&rt;</code>
tags, and then generate with `hack/update-toc.sh`.
-->

<!-- toc -->
- [Release Signoff Checklist](#release-signoff-checklist)
- [Summary](#summary)
- [Motivation](#motivation)
  - [Goals](#goals)
  - [Non-Goals](#non-goals)
- [Proposal](#proposal)
  - [User Stories (Optional)](#user-stories-optional)
    - [Story 1](#story-1)
    - [Story 2](#story-2)
  - [Notes/Constraints/Caveats (Optional)](#notesconstraintscaveats-optional)
  - [Risks and Mitigations](#risks-and-mitigations)
- [Design Details](#design-details)
  - [Test Plan](#test-plan)
      - [Prerequisite testing updates](#prerequisite-testing-updates)
      - [Unit tests](#unit-tests)
      - [Integration tests](#integration-tests)
      - [e2e tests](#e2e-tests)
  - [Graduation Criteria](#graduation-criteria)
  - [Upgrade / Downgrade Strategy](#upgrade--downgrade-strategy)
  - [Version Skew Strategy](#version-skew-strategy)
- [Production Readiness Review Questionnaire](#production-readiness-review-questionnaire)
  - [Feature Enablement and Rollback](#feature-enablement-and-rollback)
  - [Rollout, Upgrade and Rollback Planning](#rollout-upgrade-and-rollback-planning)
  - [Monitoring Requirements](#monitoring-requirements)
  - [Dependencies](#dependencies)
  - [Scalability](#scalability)
  - [Troubleshooting](#troubleshooting)
- [Implementation History](#implementation-history)
- [Drawbacks](#drawbacks)
- [Alternatives](#alternatives)
- [Infrastructure Needed (Optional)](#infrastructure-needed-optional)
<!-- /toc -->

## Release Signoff Checklist

Items marked with (R) are required *prior to targeting to a milestone / release*.

- [ ] (R) Enhancement issue in release milestone, which links to KEP dir in [kubernetes/enhancements] (not the initial KEP PR)
- [ ] (R) KEP approvers have approved the KEP status as `implementable`
- [ ] (R) Design details are appropriately documented
- [ ] (R) Test plan is in place, giving consideration to SIG Architecture and SIG Testing input (including test refactors)
  - [ ] e2e Tests for all Beta API Operations (endpoints)
  - [ ] (R) Ensure GA e2e tests meet requirements for [Conformance Tests](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/conformance-tests.md) 
  - [ ] (R) Minimum Two Week Window for GA e2e tests to prove flake free
- [ ] (R) Graduation criteria is in place
  - [ ] (R) [all GA Endpoints](https://github.com/kubernetes/community/pull/1806) must be hit by [Conformance Tests](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/conformance-tests.md) 
- [ ] (R) Production readiness review completed
- [ ] (R) Production readiness review approved
- [ ] "Implementation History" section is up-to-date for milestone
- [ ] User-facing documentation has been created in [kubernetes/website], for publication to [kubernetes.io]
- [ ] Supporting documentation—e.g., additional design documents, links to mailing list discussions/SIG meetings, relevant PRs/issues, release notes

<!--
**Note:** This checklist is iterative and should be reviewed and updated every time this enhancement is being considered for a milestone.
-->

[kubernetes.io]: https://kubernetes.io/
[kubernetes/enhancements]: https://git.k8s.io/enhancements
[kubernetes/kubernetes]: https://git.k8s.io/kubernetes
[kubernetes/website]: https://git.k8s.io/website

## Summary
<!--DIRECTIVE: required: true-->

## Motivation
<!--DIRECTIVE: required: true-->

### Goals
<!--DIRECTIVE: required: true-->

### Non-Goals
<!--DIRECTIVE: required: false-->

## Proposal
<!--DIRECTIVE: required: true-->

### User Stories (Optional)
<!--DIRECTIVE: required: false>

#### Story 1
<!--DIRECTIVE: required: false-->

#### Story 2
<!--DIRECTIVE: required: false-->

### Notes/Constraints/Caveats (Optional)
<!--DIRECTIVE: required: false>

### Risks and Mitigations
<!--DIRECTIVE: required: true-->

## Design Details
<!--DIRECTIVE: required: true-->

### Test Plan
<!--DIRECTIVE: required: true-->

[ ] I/we understand the owners of the involved components may require updates to
existing tests to make this code solid enough prior to committing the changes necessary
to implement this enhancement.

##### Prerequisite testing updates
<!--DIRECTIVE: required: false-->

##### Unit tests
<!--DIRECTIVE: required: true-->

- `<package>`: `<date>` - `<test coverage>`

##### Integration tests
<!--DIRECTIVE: required: true-->

- <test>: <link to test coverage>

##### e2e tests
<!--DIRECTIVE: required: true-->

- <test>: <link to test coverage>

### Graduation Criteria
<!--DIRECTIVE: required: true-->

### Upgrade / Downgrade Strategy
<!--DIRECTIVE: required: true-->

### Version Skew Strategy
<!--DIRECTIVE: required: true-->

## Production Readiness Review Questionnaire
<!--DIRECTIVE: required: true-->

### Feature Enablement and Rollback
<!--DIRECTIVE: required: true,stage: alpha-->

###### How can this feature be enabled / disabled in a live cluster?
<!--DIRECTIVE: required: true,stage: alpha-->

- [ ] Feature gate (also fill in values in `kep.yaml`)
  - Feature gate name:
  - Components depending on the feature gate:
- [ ] Other
  - Describe the mechanism:
  - Will enabling / disabling the feature require downtime of the control
    plane?
  - Will enabling / disabling the feature require downtime or reprovisioning
    of a node?

###### Does enabling the feature change any default behavior?
<!--DIRECTIVE: required: true,stage: alpha-->

###### Can the feature be disabled once it has been enabled (i.e. can we roll back the enablement)?
<!--DIRECTIVE: required: true,stage: alpha-->

###### What happens if we reenable the feature if it was previously rolled back?
<!--DIRECTIVE: required: true,stage: alpha-->

###### Are there any tests for feature enablement/disablement?
<!--DIRECTIVE: required: true,stage: alpha-->

### Rollout, Upgrade and Rollback Planning
<!--DIRECTIVE: required: true,stage: beta-->

###### How can a rollout or rollback fail? Can it impact already running workloads?
<!--DIRECTIVE: required: true,stage: beta-->

###### What specific metrics should inform a rollback?
<!--DIRECTIVE: required: true,stage: beta-->

###### Were upgrade and rollback tested? Was the upgrade->downgrade->upgrade path tested?
<!--DIRECTIVE: required: true,stage: beta-->

###### Is the rollout accompanied by any deprecations and/or removals of features, APIs, fields of API types, flags, etc.?
<!--DIRECTIVE: required: true,stage: beta-->

### Monitoring Requirements
<!--DIRECTIVE: required: true,stage: beta-->

###### How can an operator determine if the feature is in use by workloads?
<!--DIRECTIVE: required: true,stage: beta-->

###### How can someone using this feature know that it is working for their instance?
<!--DIRECTIVE: required: true,stage: beta-->

- [ ] Events
  - Event Reason: 
- [ ] API .status
  - Condition name: 
  - Other field: 
- [ ] Other (treat as last resort)
  - Details:

###### What are the reasonable SLOs (Service Level Objectives) for the enhancement?
<!--DIRECTIVE: required: true,stage: beta-->

###### What are the SLIs (Service Level Indicators) an operator can use to determine the health of the service?
<!--DIRECTIVE: required: true,stage: beta-->

- [ ] Metrics
  - Metric name:
  - [Optional] Aggregation method:
  - Components exposing the metric:
- [ ] Other (treat as last resort)
  - Details:

###### Are there any missing metrics that would be useful to have to improve observability of this feature?
<!--DIRECTIVE: required: true,stage: beta-->

### Dependencies
<!--DIRECTIVE: required: true,stage: beta-->

###### Does this feature depend on any specific services running in the cluster?
<!--DIRECTIVE: required: true,stage: beta-->

### Scalability
<!--DIRECTIVE: required: true,stage: beta-->

###### Will enabling / using this feature result in any new API calls?
<!--DIRECTIVE: required: true,stage: beta-->

###### Will enabling / using this feature result in introducing new API types?
<!--DIRECTIVE: required: true,stage: beta-->

###### Will enabling / using this feature result in any new calls to the cloud provider?
<!--DIRECTIVE: required: true,stage: beta-->

###### Will enabling / using this feature result in increasing size or count of the existing API objects?
<!--DIRECTIVE: required: true,stage: beta-->

###### Will enabling / using this feature result in increasing time taken by any operations covered by existing SLIs/SLOs?
<!--DIRECTIVE: required: true,stage: beta-->

###### Will enabling / using this feature result in non-negligible increase of resource usage (CPU, RAM, disk, IO, ...) in any components?
<!--DIRECTIVE: required: true,stage: beta-->

###### Can enabling / using this feature result in resource exhaustion of some node resources (PIDs, sockets, inodes, etc.)?
<!--DIRECTIVE: required: true,stage: beta-->

### Troubleshooting
<!--DIRECTIVE: required: true,stage: beta-->

###### How does this feature react if the API server and/or etcd is unavailable?
<!--DIRECTIVE: required: true,stage: beta-->

###### What are other known failure modes?
<!--DIRECTIVE: required: true,stage: beta-->

###### What steps should be taken if SLOs are not being met to determine the problem?
<!--DIRECTIVE: required: true,stage: beta-->

## Implementation History
<!--DIRECTIVE: required: true,stage: beta-->

## Drawbacks
<!--DIRECTIVE: required: true,stage: beta-->

## Alternatives
<!--DIRECTIVE: required: true-->

## Infrastructure Needed (Optional)
<!--DIRECTIVE: required: false-->
