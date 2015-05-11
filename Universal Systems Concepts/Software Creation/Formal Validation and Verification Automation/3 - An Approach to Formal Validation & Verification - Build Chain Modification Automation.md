#Build-Chain Modification Automation - An Alternative to Formal Linguistics (An Approach to Formal Validation and Verification Part 3)

##Introduction

In this third and final paper, we present an *implementation* of Formal Validation & Verification for software creation as a System Concept in practical computation.

**Build-chain** - A chain of pre-systems which converges on 'optimal compression' for some classes of System Types. 

- Formal Validator - *application* specific
	- must be constructed per application
- Formal Verifier - **build-chain** specific
	- must be constructed per build-chain (set of System Concepts that are potential dependencies)

We should not have to "code" over an absolute set of Systems Concepts when implementing a program (such as we do in a programming language).

The computer should be able to use a given Build-Chain and generate a Concept-preserving translation from System concept to implementation (with optimal compression).

If it can not generate (an optimal compression). Then it should automatically modify the Build-Chain to compensate for this defect.


##Paradox

**How can we formally validate and verify a formal validator and verifier?**

FORMAL VALIDATORS AND VERIFIERS CANNOT BE PROGRAMS! **PROGRAMS CANNOT VALIDATE AND VERIFY OTHER PROGRAMS!**

We reference Godel's Incompleteness Theorem and address the fact that Formal Linguistics (specifically programming languages) are very far left field of ever being able to address this issue in F.V.V.

Implementing a language compiler in the language that it compiles is an approach to answering this issue. We commonly call these languages: System's Programming Languages.

However, compilation is not F.V.V. and thus this cannot properly answer our initial question.









