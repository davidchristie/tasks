/// <reference path="../support/index.d.ts" />

import { v4 as uuid } from "uuid";

describe("Tasks", () => {
  beforeEach(() => {
    const userId = uuid();

    cy.request("POST", "http://localhost:8080/test/seed/reset");

    cy.request("POST", "http://localhost:8080/test/seed/users", {
      users: [
        {
          avatar_url: "AVATAR_URL",
          email: "EMAIL",
          github_id: 1,
          id: userId,
          name: "NAME",
        },
      ],
    });

    cy.request("POST", "http://localhost:8080/test/seed/tasks", {
      tasks: [
        {
          completed_at: new Date().toISOString(),
          created_at: new Date().toISOString(),
          created_by_user_id: userId,
          id: uuid(),
          text: "Task 1",
        },
        {
          completed_at: null,
          created_at: new Date().toISOString(),
          created_by_user_id: userId,
          id: uuid(),
          text: "Task 2",
        },
        {
          completed_at: null,
          created_at: new Date().toISOString(),
          created_by_user_id: userId,
          id: uuid(),
          text: "Task 3",
        },
      ],
    });

    cy.login();
  });

  it("should add new tasks to the list", () => {
    cy.getDataTest("task-list-item-text")
      .get(`input[value="New Task"]`)
      .should("not.exist");

    cy.getDataTest("add-task-button").click();

    cy.getDataTest("add-task-dialog").should("exist");

    cy.getDataTest("add-task-dialog-text").type("New Task");

    cy.getDataTest("add-task-dialog-save").click();

    cy.getDataTest("task-list-item-text")
      .get(`input[value="New Task"]`)
      .should("exist");
  });

  it("should update tasks when they have been edited", () => {
    cy.getDataTest("task-list-item-text")
      .get(`input[value="Task 2 EDIT"]`)
      .should("not.exist");

    cy.getDataTest("task-list-item-edit").first().click();

    cy.getDataTest("edit-task-dialog").should("exist");

    cy.getDataTest("edit-task-dialog-text").type(" EDIT");

    cy.getDataTest("edit-task-dialog-close").click();

    cy.getDataTest("task-list-item-text")
      .get(`input[value="Task 2 EDIT"]`)
      .should("exist");
  });

  it("should remove deleted tasks from the list", () => {
    cy.getDataTest("task-list-item-text")
      .get(`input[value="Task 2"]`)
      .should("exist");

    cy.getDataTest("task-list-item-edit").first().click();

    cy.getDataTest("edit-task-dialog").should("exist");

    cy.getDataTest("edit-task-dialog-delete").click();

    cy.getDataTest("task-list-item-text")
      .get(`input[value="Task 2"]`)
      .should("not.exist");
  });
});
