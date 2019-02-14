import React from 'react';
import {
  Card,
  Page,
  Tabs
} from '@shopify/polaris';

export default class Todo extends React.Component {
  defaultState = {
    selected: 0
  };

  state = {
    selected: this.defaultState.selected,
  };

  {% if cookiecutter.use_tabs %}
  tabs = [
    {% for tab in cookiecutter.tabs.split(" ") %}
    {
      id: "{{tab}}",
      content: "{{tab|title}}",
      accessibilityLabel: "{{tab}}",
      panelID: "{{tab}}",
    },
    {% endfor %}
  ];


  handleTabChange = (selectedTabIndex) => {
    this.setState({selected: selectedTabIndex});
  };

  getTabPanel = (tabPanels, tabIndex) => {
    let pID = this.tabs[tabIndex].panelID;
    return tabPanels[pID]
  };
  {% endif %}

  render() {
    const {selected} = this.state;

    {% if cookiecutter.use_tabs %}
    const tabPanels = {
      {%- for tab in cookiecutter.tabs.split(" ") %}
      {{tab}}: (
          <Tabs.Panel id="{{tab}}">
            <div>I'm the content for {{tab}}!</div>
          </Tabs.Panel>
        ),
      {% endfor -%}
    };
    {% endif %}


    return (
      <Page title="{{ cookiecutter.name|title}}" fullWidth>
        <Card>
          {% if cookiecutter.use_tabs %}
          <Tabs tabs={this.tabs} selected={selected} onSelect={this.handleTabChange} />
          <Card.Section title={this.tabs[selected].content}>
            {this.getTabPanel(tabPanels, selected)}
          </Card.Section>
          {% else %}
          <div>I'm the content for {{cookiecutter.name }}</div>
          {% endif %}
        </Card>
      </Page>
    );
  }
}

