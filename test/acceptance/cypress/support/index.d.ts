/// <reference types="cypress" />

declare namespace Cypress {
  interface Chainable {
    getDataTest(value: string): Chainable<Element>;
    login(): void;
  }
}
