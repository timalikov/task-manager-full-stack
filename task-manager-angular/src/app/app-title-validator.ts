import { Directive, ElementRef, HostListener, Input } from '@angular/core';

@Directive({
  selector: '[appTitleValidator]'
})
export class TitleValidatorDirective {
  @Input() appTitleValidator = 5; 

  constructor(private el: ElementRef) {}

  @HostListener('input') onInput() {
    const input = this.el.nativeElement.value;
    if (input.length < this.appTitleValidator) {
      this.el.nativeElement.style.borderColor = 'red';
    } else {
      this.el.nativeElement.style.borderColor = 'green';
    }
  }
}
