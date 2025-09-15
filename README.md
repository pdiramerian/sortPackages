Thoughtful Package Sorter

Robotic arm function to automatically dispatch packages into the correct stacks based on their volume, dimensions, and mass.

🚀 Problem Statement

The robot must decide where to send each package according to these rules:

Bulky

Volume (Width × Height × Length) >= 1,000,000 cm³ or

Any single dimension>= 150 cm

Heavy

Mass >= 20 kg

Stacks

STANDARD = Not bulky and not heavy

SPECIAL = Either bulky or heavy