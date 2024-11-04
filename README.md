# [`up` module](https://github.com/viam-modules/up)

This [up module](https://app.viam.com/module/viam/up) implements an Intel-based board like the [UP 4000](https://github.com/up-board/up-community/wiki/Pinout_UP4000) using the [`rdk:component:board` API](https://docs.viam.com/appendix/apis/components/board/).

> [!NOTE]
> Before configuring your board, you must [create a machine](https://docs.viam.com/cloud/machines/#add-a-new-machine).

## Configure your upboard board

Navigate to the [**CONFIGURE** tab](https://docs.viam.com/configure/) of your [machine](https://docs.viam.com/fleet/machines/) in the [Viam app](https://app.viam.com/).
[Add board / up:upboard to your machine](https://docs.viam.com/configure/#components).

### Attributes

The following attributes are available for `viam:up:upboard` boards:

| Attribute | Type | Required? | Description |
| --------- | ---- | --------- | ----------  |
| `digital_interrupts` | object | Optional | Any digital interrupts's pin number and name. |

For instructions on implementing digital interrupts, see [Digital interrupt configuration](#Digital-interrupt-configuration)

## Example configuration

### `viam:up:upboard`
```json
  {
    "name": "<your-up-upboard-board-name>",
    "model": "viam:up:upboard",
    "type": "board",
    "namespace": "rdk",
    "attributes": {},
    "depends_on": []
  }
```

### Next Steps
- To test your board, expand the **TEST** section of its configuration pane or go to the [**CONTROL** tab](https://docs.viam.com/fleet/control/).
- To write code against your board, use one of the [available SDKs](https://docs.viam.com/sdks/).
- To view examples using a board component, explore [these tutorials](https://docs.viam.com/tutorials/).

## Digital interrupt configuration
[Interrupts](https://en.wikipedia.org/wiki/Interrupt) are a method of signaling precise state changes.
Configuring digital interrupts to monitor GPIO pins on your board is useful when your application needs to know precisely when there is a change in GPIO value between high and low.

- When an interrupt configured on your board processes a change in the state of the GPIO pin it is configured to monitor, it ticks to record the state change.
  You can stream these ticks with the board API's [`StreamTicks()`](/appendix/apis/components/board/#streamticks), or get the current value of the digital interrupt with [`Value()`](/appendix/apis/components/board/#value).
- Calling [`GetGPIO()`](/appendix/apis/components/board/#getgpio) on a GPIO pin, which you can do without configuring interrupts, is useful when you want to know a pin's value at specific points in your program, but is less precise and convenient than using an interrupt.

Integrate `digital_interrupts` into your machine in the `attributes` of your board by adding the following to your board's `attributes` configuration:

```json {class="line-numbers linkable-line-numbers"}
{
  "digital_interrupts": [
    {
      "name": "<your-digital-interrupt-name>",
      "pin": "<your-digital-interrupt-pin-number>"
    }
  ]
}
```

### Attributes

The following attributes are available for `digital_interrupts`:

<!-- prettier-ignore -->
| Name | Type | Required? | Description |
| ---- | ---- | --------- | ----------- |
|`name` | string | **Required** | Your name for the digital interrupt. |
|`pin`| string | **Required** | The pin number of the board's GPIO pin that you wish to configure the digital interrupt for. |

### Example configuration

```json {class="line-numbers linkable-line-numbers"}
{
  "components": [
    {
      "name": "<your-up-upboard-board-name>",
      "model": "viam:up:upboard",
      "type": "board",
      "namespace": "rdk",
      "attributes": {
        "digital_interrupts": [
          {
            "name": "your-interrupt-1",
            "pin": "15"
          },
          {
            "name": "your-interrupt-2",
            "pin": "16"
          }
        ]
      }
    }
  ]
}
```
